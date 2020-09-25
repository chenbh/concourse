package present

import (
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/types"
)

func Check(check db.Check) types.Check {

	atcCheck := types.Check{
		ID:     check.ID(),
		Status: string(check.Status()),
	}

	if !check.CreateTime().IsZero() {
		atcCheck.CreateTime = check.CreateTime().Unix()
	}

	if !check.StartTime().IsZero() {
		atcCheck.StartTime = check.StartTime().Unix()
	}

	if !check.EndTime().IsZero() {
		atcCheck.EndTime = check.EndTime().Unix()
	}

	if err := check.CheckError(); err != nil {
		atcCheck.CheckError = err.Error()
	}

	return atcCheck
}
