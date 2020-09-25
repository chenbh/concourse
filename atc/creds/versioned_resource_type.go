package creds

import (
	"github.com/concourse/concourse/atc/types"
	"github.com/concourse/concourse/vars"
)

type VersionedResourceType struct {
	types.VersionedResourceType

	Source Source
}

type VersionedResourceTypes []VersionedResourceType

func NewVersionedResourceTypes(variables vars.Variables, rawTypes types.VersionedResourceTypes) VersionedResourceTypes {
	var types VersionedResourceTypes
	for _, t := range rawTypes {
		types = append(types, VersionedResourceType{
			VersionedResourceType: t,
			Source:                NewSource(variables, t.Source),
		})
	}

	return types
}

func (vrt VersionedResourceTypes) Lookup(name string) (VersionedResourceType, bool) {
	for _, t := range vrt {
		if t.Name == name {
			return t, true
		}
	}

	return VersionedResourceType{}, false
}

func (vrt VersionedResourceTypes) Without(name string) VersionedResourceTypes {
	newTypes := VersionedResourceTypes{}

	for _, t := range vrt {
		if t.Name != name {
			newTypes = append(newTypes, t)
		}
	}

	return newTypes
}

func (vrt VersionedResourceTypes) Evaluate() (types.VersionedResourceTypes, error) {

	var rawTypes types.VersionedResourceTypes
	for _, t := range vrt {
		source, err := t.Source.Evaluate()
		if err != nil {
			return nil, err
		}

		resourceType := t.ResourceType
		resourceType.Source = source

		rawTypes = append(rawTypes, types.VersionedResourceType{
			ResourceType: resourceType,
			Version:      t.Version,
		})
	}

	return rawTypes, nil
}
