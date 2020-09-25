package types

import (
	"fmt"
	"sigs.k8s.io/yaml"
)

const ConfigVersionHeader = "X-Concourse-Config-Version"
const DefaultTeamName = "main"

type Tags []string

type Config struct {
	Groups        GroupConfigs     `json:"groups,omitempty"`
	VarSources    VarSourceConfigs `json:"var_sources,omitempty"`
	Resources     ResourceConfigs  `json:"resources,omitempty"`
	ResourceTypes ResourceTypes    `json:"resource_types,omitempty"`
	Jobs          JobConfigs       `json:"jobs,omitempty"`
	Display       *DisplayConfig   `json:"display,omitempty"`
}

func UnmarshalConfig(payload []byte, config interface{}) error {
	// a 'skeleton' of Config, specifying only the toplevel fields
	type skeletonConfig struct {
		Groups        interface{} `json:"groups,omitempty"`
		VarSources    interface{} `json:"var_sources,omitempty"`
		Resources     interface{} `json:"resources,omitempty"`
		ResourceTypes interface{} `json:"resource_types,omitempty"`
		Jobs          interface{} `json:"jobs,omitempty"`
		Display       interface{} `json:"display,omitempty"`
	}

	var stripped skeletonConfig
	err := yaml.Unmarshal(payload, &stripped)
	if err != nil {
		return err
	}

	strippedPayload, err := yaml.Marshal(stripped)
	if err != nil {
		return err
	}

	return yaml.UnmarshalStrict(
		strippedPayload,
		&config,
	)
}

type GroupConfig struct {
	Name      string   `json:"name"`
	Jobs      []string `json:"jobs,omitempty"`
	Resources []string `json:"resources,omitempty"`
}

type GroupConfigs []GroupConfig

func (groups GroupConfigs) Lookup(name string) (GroupConfig, int, bool) {
	for index, group := range groups {
		if group.Name == name {
			return group, index, true
		}
	}

	return GroupConfig{}, -1, false
}

type VarSourceConfig struct {
	Name   string      `json:"name"`
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
}

type VarSourceConfigs []VarSourceConfig

func (c VarSourceConfigs) Lookup(name string) (VarSourceConfig, bool) {
	for _, cm := range c {
		if cm.Name == name {
			return cm, true
		}
	}

	return VarSourceConfig{}, false
}

type ResourceConfig struct {
	Name         string  `json:"name"`
	OldName      string  `json:"old_name,omitempty"`
	Public       bool    `json:"public,omitempty"`
	WebhookToken string  `json:"webhook_token,omitempty"`
	Type         string  `json:"type"`
	Source       Source  `json:"source"`
	CheckEvery   string  `json:"check_every,omitempty"`
	CheckTimeout string  `json:"check_timeout,omitempty"`
	Tags         Tags    `json:"tags,omitempty"`
	Version      Version `json:"version,omitempty"`
	Icon         string  `json:"icon,omitempty"`
}

type ResourceType struct {
	Name                 string `json:"name"`
	Type                 string `json:"type"`
	Source               Source `json:"source"`
	Privileged           bool   `json:"privileged,omitempty"`
	CheckEvery           string `json:"check_every,omitempty"`
	Tags                 Tags   `json:"tags,omitempty"`
	Params               Params `json:"params,omitempty"`
	CheckSetupError      string `json:"check_setup_error,omitempty"`
	CheckError           string `json:"check_error,omitempty"`
	UniqueVersionHistory bool   `json:"unique_version_history,omitempty"`
}

type DisplayConfig struct {
	BackgroundImage string `json:"background_image,omitempty"`
}

type ResourceTypes []ResourceType

func (types ResourceTypes) Lookup(name string) (ResourceType, bool) {
	for _, t := range types {
		if t.Name == name {
			return t, true
		}
	}

	return ResourceType{}, false
}

func (types ResourceTypes) Without(name string) ResourceTypes {
	newTypes := ResourceTypes{}
	for _, t := range types {
		if t.Name != name {
			newTypes = append(newTypes, t)
		}
	}

	return newTypes
}

type ResourceConfigs []ResourceConfig

func (resources ResourceConfigs) Lookup(name string) (ResourceConfig, bool) {
	for _, resource := range resources {
		if resource.Name == name {
			return resource, true
		}
	}

	return ResourceConfig{}, false
}

type JobConfigs []JobConfig

func (jobs JobConfigs) Lookup(name string) (JobConfig, bool) {
	for _, job := range jobs {
		if job.Name == name {
			return job, true
		}
	}

	return JobConfig{}, false
}

func (config Config) JobIsPublic(jobName string) (bool, error) {
	job, found := config.Jobs.Lookup(jobName)
	if !found {
		return false, fmt.Errorf("cannot find job with job name '%s'", jobName)
	}

	return job.Public, nil
}

