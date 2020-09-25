package event

import (
	"github.com/concourse/concourse/atc/types"
)

const (
	// build log (e.g. from input or build execution)
	EventTypeLog types.EventType = "log"

	// build status change (e.g. 'started', 'succeeded')
	EventTypeStatus types.EventType = "status"

	// a step (get/put/task) selected worker
	EventTypeSelectedWorker types.EventType = "selected-worker"

	// task execution started
	EventTypeStartTask types.EventType = "start-task"

	// task initializing (all inputs fetched; fetching image)
	EventTypeInitializeTask types.EventType = "initialize-task"

	// task execution finished
	EventTypeFinishTask types.EventType = "finish-task"

	// initialize getting something
	EventTypeInitializeGet types.EventType = "initialize-get"

	// started getting something
	EventTypeStartGet types.EventType = "start-get"

	// finished getting something
	EventTypeFinishGet types.EventType = "finish-get"

	// initialize putting something
	EventTypeInitializePut types.EventType = "initialize-put"

	// started putting something
	EventTypeStartPut types.EventType = "start-put"

	// finished putting something
	EventTypeFinishPut types.EventType = "finish-put"

	EventTypeSetPipelineChanged types.EventType = "set-pipeline-changed"

	// initialize step
	EventTypeInitialize types.EventType = "initialize"

	// started step
	EventTypeStart types.EventType = "start"

	// finished step
	EventTypeFinish types.EventType = "finish"

	// error occurred
	EventTypeError types.EventType = "error"
)
