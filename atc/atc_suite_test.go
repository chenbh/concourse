package atc_test

import (
	"github.com/concourse/concourse/atc/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"testing"
)

func TestATC(t *testing.T) {
	suite.Run(t, &types.StepsSuite{
		Assertions: require.New(t),
	})

	RegisterFailHandler(Fail)
	RunSpecs(t, "ATC Suite")
}
