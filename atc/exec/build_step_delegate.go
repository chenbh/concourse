package exec

import (
	"io"

	"code.cloudfoundry.org/lager"

	"github.com/chenbh/concourse/atc"
	"github.com/chenbh/concourse/atc/db"
	"github.com/chenbh/concourse/vars"
)

//go:generate counterfeiter . BuildStepDelegate

type BuildStepDelegate interface {
	ImageVersionDetermined(db.UsedResourceCache) error
	RedactImageSource(source atc.Source) (atc.Source, error)

	Stdout() io.Writer
	Stderr() io.Writer

	Variables() vars.CredVarsTracker

	Initializing(lager.Logger)
	Starting(lager.Logger)
	Finished(lager.Logger, bool)
	SelectedWorker(lager.Logger, string)
	Errored(lager.Logger, string)
}
