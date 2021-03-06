package db

import (
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/concourse/concourse/atc/types"

	sq "github.com/Masterminds/squirrel"
	"github.com/concourse/concourse/atc/db/lock"
)

var ErrResourceCacheAlreadyExists = errors.New("resource-cache-already-exists")
var ErrResourceCacheDisappeared = errors.New("resource-cache-disappeared")

// ResourceCache represents an instance of a ResourceConfig's version.
//
// A ResourceCache is created by a `get`, an `image_resource`, or a resource
// type in a pipeline.
//
// ResourceCaches are garbage-collected by gc.ResourceCacheCollector.
type ResourceCacheDescriptor struct {
	ResourceConfigDescriptor ResourceConfigDescriptor // The resource configuration.
	Version                  types.Version            // The version of the resource.
	Params                   types.Params             // The params used when fetching the version.
}

func (cache *ResourceCacheDescriptor) find(tx Tx, lockFactory lock.LockFactory, conn Conn) (UsedResourceCache, bool, error) {
	resourceConfig, found, err := cache.ResourceConfigDescriptor.find(tx, lockFactory, conn)
	if err != nil {
		return nil, false, err
	}

	if !found {
		return nil, false, nil
	}

	return cache.findWithResourceConfig(tx, resourceConfig, lockFactory, conn)
}

func (cache *ResourceCacheDescriptor) findOrCreate(
	tx Tx,
	lockFactory lock.LockFactory,
	conn Conn,
) (UsedResourceCache, error) {
	resourceConfig, err := cache.ResourceConfigDescriptor.findOrCreate(tx, lockFactory, conn)
	if err != nil {
		return nil, err
	}

	rc, found, err := cache.findWithResourceConfig(tx, resourceConfig, lockFactory, conn)
	if err != nil {
		return nil, err
	}

	if !found {
		var id int
		err = psql.Insert("resource_caches").
			Columns(
				"resource_config_id",
				"version",
				"version_md5",
				"params_hash",
			).
			Values(
				resourceConfig.ID(),
				cache.version(),
				sq.Expr("md5(?)", cache.version()),
				paramsHash(cache.Params),
			).
			Suffix(`
				ON CONFLICT (resource_config_id, version_md5, params_hash) DO UPDATE SET
				resource_config_id = EXCLUDED.resource_config_id,
				version = EXCLUDED.version,
				version_md5 = EXCLUDED.version_md5,
				params_hash = EXCLUDED.params_hash
				RETURNING id
			`).
			RunWith(tx).
			QueryRow().
			Scan(&id)
		if err != nil {
			return nil, err
		}

		rc = &usedResourceCache{
			id:             id,
			version:        cache.Version,
			resourceConfig: resourceConfig,
			lockFactory:    lockFactory,
			conn:           conn,
		}
	}

	return rc, nil
}

func (cache *ResourceCacheDescriptor) use(
	tx Tx,
	rc UsedResourceCache,
	user ResourceCacheUser,
) error {
	cols := user.SQLMap()
	cols["resource_cache_id"] = rc.ID()

	var resourceCacheUseExists int
	err := psql.Select("1").
		From("resource_cache_uses").
		Where(sq.Eq(cols)).
		RunWith(tx).
		QueryRow().
		Scan(&resourceCacheUseExists)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}

	if err == nil {
		// use already exists
		return nil
	}

	_, err = psql.Insert("resource_cache_uses").
		SetMap(cols).
		RunWith(tx).
		Exec()
	return err
}

func (cache *ResourceCacheDescriptor) findWithResourceConfig(tx Tx, resourceConfig ResourceConfig, lockFactory lock.LockFactory, conn Conn) (UsedResourceCache, bool, error) {
	var id int
	err := psql.Select("id").
		From("resource_caches").
		Where(sq.Eq{
			"resource_config_id": resourceConfig.ID(),
			"params_hash":        paramsHash(cache.Params),
		}).
		Where(sq.Expr("version_md5 = md5(?)", cache.version())).
		Suffix("FOR SHARE").
		RunWith(tx).
		QueryRow().
		Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}

		return nil, false, err
	}

	return &usedResourceCache{
		id:             id,
		version:        cache.Version,
		resourceConfig: resourceConfig,
		lockFactory:    lockFactory,
		conn:           conn,
	}, true, nil
}

func (cache *ResourceCacheDescriptor) version() string {
	j, _ := json.Marshal(cache.Version)
	return string(j)
}

func paramsHash(p types.Params) string {
	if p != nil {
		return mapHash(p)
	}

	return mapHash(types.Params{})
}

// UsedResourceCache is created whenever a ResourceCache is Created and/or
// Used.
//
// So long as the UsedResourceCache exists, the underlying ResourceCache can
// not be removed.
//
// UsedResourceCaches become unused by the gc.ResourceCacheCollector, which may
// then lead to the ResourceCache being garbage-collected.
//
// See FindOrCreateForBuild, FindOrCreateForResource, and
// FindOrCreateForResourceType for more information on when it becomes unused.

//go:generate counterfeiter . UsedResourceCache

type UsedResourceCache interface {
	ID() int
	Version() types.Version

	ResourceConfig() ResourceConfig

	Destroy(Tx) (bool, error)
	BaseResourceType() *UsedBaseResourceType
}

type usedResourceCache struct {
	id             int
	resourceConfig ResourceConfig
	version        types.Version

	lockFactory lock.LockFactory
	conn        Conn
}

func (cache *usedResourceCache) ID() int                        { return cache.id }
func (cache *usedResourceCache) ResourceConfig() ResourceConfig { return cache.resourceConfig }
func (cache *usedResourceCache) Version() types.Version         { return cache.version }

func (cache *usedResourceCache) Destroy(tx Tx) (bool, error) {
	rows, err := psql.Delete("resource_caches").
		Where(sq.Eq{
			"id": cache.id,
		}).
		RunWith(tx).
		Exec()
	if err != nil {
		return false, err
	}

	affected, err := rows.RowsAffected()
	if err != nil {
		return false, err
	}

	if affected == 0 {
		return false, ErrResourceCacheDisappeared
	}

	return true, nil
}

func (cache *usedResourceCache) BaseResourceType() *UsedBaseResourceType {
	if cache.resourceConfig.CreatedByBaseResourceType() != nil {
		return cache.resourceConfig.CreatedByBaseResourceType()
	}

	return cache.resourceConfig.CreatedByResourceCache().BaseResourceType()
}

func mapHash(m map[string]interface{}) string {
	j, _ := json.Marshal(m)
	return fmt.Sprintf("%x", sha256.Sum256(j))
}
