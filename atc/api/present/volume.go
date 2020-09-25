package present

import (
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/types"
)

func Volume(volume db.CreatedVolume) (types.Volume, error) {
	resourceType, err := volume.ResourceType()
	if err != nil {
		return types.Volume{}, err
	}

	baseResourceType, err := volume.BaseResourceType()
	if err != nil {
		return types.Volume{}, err
	}

	pipelineName, jobName, stepName, err := volume.TaskIdentifier()
	if err != nil {
		return types.Volume{}, err
	}

	return types.Volume{
		ID:               volume.Handle(),
		Type:             string(volume.Type()),
		WorkerName:       volume.WorkerName(),
		ContainerHandle:  volume.ContainerHandle(),
		Path:             volume.Path(),
		ParentHandle:     volume.ParentHandle(),
		PipelineName:     pipelineName,
		JobName:          jobName,
		StepName:         stepName,
		ResourceType:     toVolumeResourceType(resourceType),
		BaseResourceType: toVolumeBaseResourceType(baseResourceType),
	}, nil
}

func toVolumeResourceType(dbResourceType *db.VolumeResourceType) *types.VolumeResourceType {
	if dbResourceType == nil {
		return nil
	}

	if dbResourceType.WorkerBaseResourceType != nil {
		return &types.VolumeResourceType{
			BaseResourceType: toVolumeBaseResourceType(dbResourceType.WorkerBaseResourceType),
			Version:          dbResourceType.Version,
		}
	}

	if dbResourceType.ResourceType != nil {
		resourceType := toVolumeResourceType(dbResourceType.ResourceType)
		return &types.VolumeResourceType{
			ResourceType: resourceType,
			Version:      dbResourceType.Version,
		}
	}

	return nil
}

func toVolumeBaseResourceType(dbResourceType *db.UsedWorkerBaseResourceType) *types.VolumeBaseResourceType {
	if dbResourceType == nil {
		return nil
	}

	return &types.VolumeBaseResourceType{
		Name:    dbResourceType.Name,
		Version: dbResourceType.Version,
	}
}
