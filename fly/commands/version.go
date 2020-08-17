package commands

import (
	"fmt"
	"os"

	"github.com/chenbh/concourse/v6"
)

func init() {
	Fly.Version = func() {
		fmt.Println(concourse.Version)
		os.Exit(0)
	}
}
