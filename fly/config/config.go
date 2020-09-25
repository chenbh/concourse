package config

import (
	"github.com/concourse/concourse/atc/types"
	"syscall"
)

func OverrideTaskParams(configFile []byte, args []string) (types.TaskConfig, error) {
	config, err := types.NewTaskConfig(configFile)
	if err != nil {
		return types.TaskConfig{}, err
	}

	config.Run.Args = append(config.Run.Args, args...)

	for k := range config.Params {
		env, found := syscall.Getenv(k)
		if found {
			config.Params[k] = env
		}
	}

	return config, nil
}
