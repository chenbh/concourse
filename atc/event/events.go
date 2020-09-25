package event

import (
	"github.com/concourse/concourse/atc/types"
)

type Error struct {
	Message string `json:"message"`
	Origin  Origin `json:"origin"`
	Time    int64  `json:"time"`
}

func (Error) EventType() types.EventType  { return EventTypeError }
func (Error) Version() types.EventVersion { return "4.1" }

type FinishTask struct {
	Time       int64  `json:"time"`
	ExitStatus int    `json:"exit_status"`
	Origin     Origin `json:"origin"`
}

func (FinishTask) EventType() types.EventType  { return EventTypeFinishTask }
func (FinishTask) Version() types.EventVersion { return "4.0" }

type InitializeTask struct {
	Time       int64      `json:"time"`
	Origin     Origin     `json:"origin"`
	TaskConfig TaskConfig `json:"config"`
}

func (InitializeTask) EventType() types.EventType  { return EventTypeInitializeTask }
func (InitializeTask) Version() types.EventVersion { return "4.0" }

// shadow the real atc.TaskConfig
type TaskConfig struct {
	Platform string `json:"platform"`
	Image    string `json:"image"`

	Run    TaskRunConfig     `json:"run"`
	Inputs []TaskInputConfig `json:"inputs"`
}

type TaskRunConfig struct {
	Path string   `json:"path"`
	Args []string `json:"args"`
	Dir  string   `json:"dir"`
}

type TaskInputConfig struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func ShadowTaskConfig(config types.TaskConfig) TaskConfig {
	var inputConfigs []TaskInputConfig

	for _, input := range config.Inputs {
		inputConfigs = append(inputConfigs, TaskInputConfig{
			Name: input.Name,
			Path: input.Path,
		})
	}

	return TaskConfig{
		Platform: config.Platform,
		Image:    config.RootfsURI,
		Run: TaskRunConfig{
			Path: config.Run.Path,
			Args: config.Run.Args,
			Dir:  config.Run.Dir,
		},
		Inputs: inputConfigs,
	}
}

type StartTask struct {
	Time       int64      `json:"time"`
	Origin     Origin     `json:"origin"`
	TaskConfig TaskConfig `json:"config"`
}

func (StartTask) EventType() types.EventType  { return EventTypeStartTask }
func (StartTask) Version() types.EventVersion { return "5.0" }

type Status struct {
	Status types.BuildStatus `json:"status"`
	Time   int64             `json:"time"`
}

func (Status) EventType() types.EventType  { return EventTypeStatus }
func (Status) Version() types.EventVersion { return "1.0" }

type SelectedWorker struct {
	Time       int64  `json:"time"`
	Origin     Origin `json:"origin"`
	WorkerName string `json:"selected_worker"`
}

func (SelectedWorker) EventType() types.EventType  { return EventTypeSelectedWorker }
func (SelectedWorker) Version() types.EventVersion { return "1.0" }

type Log struct {
	Time    int64  `json:"time"`
	Origin  Origin `json:"origin"`
	Payload string `json:"payload"`
}

func (Log) EventType() types.EventType  { return EventTypeLog }
func (Log) Version() types.EventVersion { return "5.1" }

type Origin struct {
	ID     OriginID     `json:"id,omitempty"`
	Source OriginSource `json:"source,omitempty"`
}

type OriginID string

type OriginSource string

const (
	OriginSourceStdout OriginSource = "stdout"
	OriginSourceStderr OriginSource = "stderr"
)

type InitializeGet struct {
	Origin Origin `json:"origin"`
	Time   int64  `json:"time,omitempty"`
}

func (InitializeGet) EventType() types.EventType  { return EventTypeInitializeGet }
func (InitializeGet) Version() types.EventVersion { return "2.0" }

type StartGet struct {
	Origin Origin `json:"origin"`
	Time   int64  `json:"time,omitempty"`
}

func (StartGet) EventType() types.EventType  { return EventTypeStartGet }
func (StartGet) Version() types.EventVersion { return "1.0" }

type FinishGet struct {
	Origin          Origin                `json:"origin"`
	Time            int64                 `json:"time"`
	ExitStatus      int                   `json:"exit_status"`
	FetchedVersion  types.Version         `json:"version"`
	FetchedMetadata []types.MetadataField `json:"metadata,omitempty"`
}

func (FinishGet) EventType() types.EventType  { return EventTypeFinishGet }
func (FinishGet) Version() types.EventVersion { return "5.1" }

type InitializePut struct {
	Origin Origin `json:"origin"`
	Time   int64  `json:"time,omitempty"`
}

func (InitializePut) EventType() types.EventType  { return EventTypeInitializePut }
func (InitializePut) Version() types.EventVersion { return "2.0" }

type StartPut struct {
	Origin Origin `json:"origin"`
	Time   int64  `json:"time,omitempty"`
}

func (StartPut) EventType() types.EventType  { return EventTypeStartPut }
func (StartPut) Version() types.EventVersion { return "1.0" }

type FinishPut struct {
	Origin          Origin                `json:"origin"`
	Time            int64                 `json:"time"`
	ExitStatus      int                   `json:"exit_status"`
	CreatedVersion  types.Version         `json:"version"`
	CreatedMetadata []types.MetadataField `json:"metadata,omitempty"`
}

func (FinishPut) EventType() types.EventType  { return EventTypeFinishPut }
func (FinishPut) Version() types.EventVersion { return "5.1" }

type SetPipelineChanged struct {
	Origin  Origin `json:"origin"`
	Changed bool   `json:"changed"`
}

func (SetPipelineChanged) EventType() types.EventType  { return EventTypeSetPipelineChanged }
func (SetPipelineChanged) Version() types.EventVersion { return "1.0" }

type Initialize struct {
	Origin Origin `json:"origin"`
	Time   int64  `json:"time,omitempty"`
}

func (Initialize) EventType() types.EventType  { return EventTypeInitialize }
func (Initialize) Version() types.EventVersion { return "1.0" }

type Start struct {
	Origin Origin `json:"origin"`
	Time   int64  `json:"time,omitempty"`
}

func (Start) EventType() types.EventType  { return EventTypeStart }
func (Start) Version() types.EventVersion { return "1.0" }

type Finish struct {
	Origin    Origin `json:"origin"`
	Time      int64  `json:"time"`
	Succeeded bool   `json:"succeeded"`
}

func (Finish) EventType() types.EventType  { return EventTypeFinish }
func (Finish) Version() types.EventVersion { return "1.0" }
