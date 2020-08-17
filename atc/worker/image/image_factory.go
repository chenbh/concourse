package image

import (
	"errors"

	"code.cloudfoundry.org/lager"
	"github.com/chenbh/concourse/v6/atc"
	"github.com/chenbh/concourse/v6/atc/compression"
	"github.com/chenbh/concourse/v6/atc/worker"
	w "github.com/chenbh/concourse/v6/atc/worker"
)

var ErrUnsupportedResourceType = errors.New("unsupported resource type")

type imageFactory struct {
	imageResourceFetcherFactory ImageResourceFetcherFactory
	compression                 compression.Compression
}

func NewImageFactory(
	imageResourceFetcherFactory ImageResourceFetcherFactory,
	compression compression.Compression,
) worker.ImageFactory {
	return &imageFactory{
		imageResourceFetcherFactory: imageResourceFetcherFactory,
		compression:                 compression,
	}
}

func (f *imageFactory) GetImage(
	logger lager.Logger,
	worker worker.Worker,
	volumeClient worker.VolumeClient,
	imageSpec worker.ImageSpec,
	teamID int,
	delegate worker.ImageFetchingDelegate,
	resourceTypes atc.VersionedResourceTypes,
) (worker.Image, error) {
	if imageSpec.ImageArtifactSource != nil {
		artifactVolume, existsOnWorker, err := imageSpec.ImageArtifactSource.ExistsOn(logger, worker)
		if err != nil {
			logger.Error("failed-to-check-if-volume-exists-on-worker", err)
			return nil, err
		}

		if existsOnWorker {
			return &imageProvidedByPreviousStepOnSameWorker{
				artifactVolume: artifactVolume,
				imageSpec:      imageSpec,
				teamID:         teamID,
				volumeClient:   volumeClient,
			}, nil
		}

		return &imageProvidedByPreviousStepOnDifferentWorker{
			imageSpec:    imageSpec,
			teamID:       teamID,
			volumeClient: volumeClient,
		}, nil
	}

	// check if custom resource
	resourceType, found := resourceTypes.Lookup(imageSpec.ResourceType)
	if found {
		imageResourceFetcher := f.imageResourceFetcherFactory.NewImageResourceFetcher(
			worker,
			w.ImageResource{
				Type:   resourceType.Type,
				Source: resourceType.Source,
				Params: resourceType.Params,
			},
			resourceType.Version,
			teamID,
			resourceTypes.Without(imageSpec.ResourceType),
			delegate,
			f.compression,
		)

		return &imageFromResource{
			imageResourceFetcher: imageResourceFetcher,

			privileged:   resourceType.Privileged,
			teamID:       teamID,
			volumeClient: volumeClient,
		}, nil
	}

	if imageSpec.ImageResource != nil {
		var version atc.Version
		version = imageSpec.ImageResource.Version

		imageResourceFetcher := f.imageResourceFetcherFactory.NewImageResourceFetcher(
			worker,
			*imageSpec.ImageResource,
			version,
			teamID,
			resourceTypes,
			delegate,
			f.compression,
		)

		return &imageFromResource{
			imageResourceFetcher: imageResourceFetcher,

			privileged:   imageSpec.Privileged,
			teamID:       teamID,
			volumeClient: volumeClient,
		}, nil
	}

	if imageSpec.ResourceType != "" {
		return &imageFromBaseResourceType{
			worker:           worker,
			resourceTypeName: imageSpec.ResourceType,
			teamID:           teamID,
			volumeClient:     volumeClient,
		}, nil
	}

	return &imageFromRootfsURI{
		url: imageSpec.ImageURL,
	}, nil
}
