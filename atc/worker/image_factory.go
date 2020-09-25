package worker

import (
	"code.cloudfoundry.org/lager"
	"context"
	"github.com/concourse/concourse/atc/types"
	"io"
	"io/ioutil"

	"github.com/concourse/concourse/atc/db"
)

//go:generate counterfeiter . ImageFactory

type ImageFactory interface {
	GetImage(
		logger lager.Logger,
		worker Worker,
		volumeClient VolumeClient,
		imageSpec ImageSpec,
		teamID int,
		delegate ImageFetchingDelegate,
		resourceTypes types.VersionedResourceTypes,
	) (Image, error)
}

type FetchedImage struct {
	Metadata   ImageMetadata
	Version    types.Version
	URL        string
	Privileged bool
}

//go:generate counterfeiter . Image

type Image interface {
	FetchForContainer(
		ctx context.Context,
		logger lager.Logger,
		container db.CreatingContainer,
	) (FetchedImage, error)
}

//go:generate counterfeiter . ImageFetchingDelegate

type ImageFetchingDelegate interface {
	Stdout() io.Writer
	Stderr() io.Writer
	ImageVersionDetermined(db.UsedResourceCache) error

	RedactImageSource(source types.Source) (types.Source, error)
}

type ImageMetadata struct {
	Env  []string `json:"env"`
	User string   `json:"user"`
}

type NoopImageFetchingDelegate struct{}

func (NoopImageFetchingDelegate) Stdout() io.Writer                                 { return ioutil.Discard }
func (NoopImageFetchingDelegate) Stderr() io.Writer                                 { return ioutil.Discard }
func (NoopImageFetchingDelegate) ImageVersionDetermined(db.UsedResourceCache) error { return nil }
func (NoopImageFetchingDelegate) RedactImageSource(source types.Source) (types.Source, error) {
	// As this is noop, redaction can just return an empty source.
	return types.Source{}, nil
}
