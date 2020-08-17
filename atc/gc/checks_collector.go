package gc

import (
	"context"
	"time"

	"code.cloudfoundry.org/lager/lagerctx"
	"github.com/chenbh/concourse/v6/atc/db"
	"github.com/chenbh/concourse/v6/atc/metric"
)

type checkCollector struct {
	checkLifecycle db.CheckLifecycle
	recyclePeriod  time.Duration
}

func NewCheckCollector(checkLifecycle db.CheckLifecycle, recyclePeriod time.Duration) *checkCollector {
	return &checkCollector{
		checkLifecycle: checkLifecycle,
		recyclePeriod:  recyclePeriod,
	}
}

func (c *checkCollector) Run(ctx context.Context) error {
	logger := lagerctx.FromContext(ctx).Session("check-collector")

	logger.Debug("start")
	defer logger.Debug("done")

	deleted, err := c.checkLifecycle.RemoveExpiredChecks(c.recyclePeriod)
	if err != nil {
		logger.Error("failed-to-remove-expired-checks", err)
		return err
	}

	metric.ChecksDeleted.IncDelta(deleted)

	return nil
}
