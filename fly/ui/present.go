package ui

import (
	"github.com/concourse/concourse/atc/types"
	"sort"
	"strings"
)

func PresentVersion(version types.Version) string {
	pairs := []string{}
	for k, v := range version {
		pairs = append(pairs, k+":"+v)
	}

	// consistent ordering
	sort.Strings(pairs)

	return strings.Join(pairs, ",")
}
