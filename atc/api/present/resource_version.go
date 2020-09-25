package present

import (
	"github.com/concourse/concourse/atc/types"
)

func ResourceVersions(hideMetadata bool, resourceVersions []types.ResourceVersion) []types.ResourceVersion {
	presented := []types.ResourceVersion{}

	for _, resourceVersion := range resourceVersions {
		if hideMetadata {
			resourceVersion.Metadata = nil
		}

		presented = append(presented, resourceVersion)
	}

	return presented
}
