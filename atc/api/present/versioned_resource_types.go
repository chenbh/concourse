package present

import (
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/types"
)

func VersionedResourceTypes(showCheckError bool, savedResourceTypes db.ResourceTypes) types.VersionedResourceTypes {
	versionedResourceTypes := savedResourceTypes.Deserialize()

	for i, resourceType := range savedResourceTypes {
		if resourceType.CheckSetupError() != nil && showCheckError {
			versionedResourceTypes[i].CheckSetupError = resourceType.CheckSetupError().Error()
		} else {
			versionedResourceTypes[i].CheckSetupError = ""
		}

		if resourceType.CheckError() != nil && showCheckError {
			versionedResourceTypes[i].CheckError = resourceType.CheckError().Error()
		} else {
			versionedResourceTypes[i].CheckError = ""
		}
	}

	return versionedResourceTypes
}
