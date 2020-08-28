package commands

import (
	"fmt"
	"os"

	"github.com/chenbh/concourse"
)

func init() {
	Fly.Version = func() {
		fmt.Println(concourse.Version)
		os.Exit(0)
	}
}
