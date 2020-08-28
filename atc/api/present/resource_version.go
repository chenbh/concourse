package present

import (
	"github.com/chenbh/concourse/atc"
)

func ResourceVersions(hideMetadata bool, resourceVersions []atc.ResourceVersion) []atc.ResourceVersion {
	presented := []atc.ResourceVersion{}

	for _, resourceVersion := range resourceVersions {
		if hideMetadata {
			resourceVersion.Metadata = nil
		}

		presented = append(presented, resourceVersion)
	}

	return presented
}
