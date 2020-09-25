package event

import (
	"encoding/json"
	"fmt"
	"github.com/concourse/concourse/atc/types"
	"reflect"
	"strings"
)

type eventTable map[types.EventType]eventVersions
type eventVersions map[types.EventVersion]eventParser
type eventParser func([]byte) (types.Event, error)

var events = eventTable{}

func unmarshaler(e types.Event) func([]byte) (types.Event, error) {
	return func(payload []byte) (types.Event, error) {
		val := reflect.New(reflect.TypeOf(e))
		err := json.Unmarshal(payload, val.Interface())
		return val.Elem().Interface().(types.Event), err
	}
}

func RegisterEvent(e types.Event) {
	versions, found := events[e.EventType()]
	if !found {
		versions = eventVersions{}
		events[e.EventType()] = versions
	}

	versions[e.Version()] = unmarshaler(e)
}

func init() {
	RegisterEvent(InitializeTask{})
	RegisterEvent(StartTask{})
	RegisterEvent(FinishTask{})
	RegisterEvent(InitializeGet{})
	RegisterEvent(StartGet{})
	RegisterEvent(FinishGet{})
	RegisterEvent(InitializePut{})
	RegisterEvent(StartPut{})
	RegisterEvent(FinishPut{})
	RegisterEvent(SetPipelineChanged{})
	RegisterEvent(Status{})
	RegisterEvent(SelectedWorker{})
	RegisterEvent(Log{})
	RegisterEvent(Error{})

	// deprecated:
	RegisterEvent(InitializeV10{})
	RegisterEvent(FinishV10{})
	RegisterEvent(StartV10{})
	RegisterEvent(InputV10{})
	RegisterEvent(InputV20{})
	RegisterEvent(OutputV10{})
	RegisterEvent(OutputV20{})
	RegisterEvent(ErrorV10{})
	RegisterEvent(ErrorV20{})
	RegisterEvent(ErrorV30{})
	RegisterEvent(FinishTaskV10{})
	RegisterEvent(FinishTaskV20{})
	RegisterEvent(FinishTaskV30{})
	RegisterEvent(InitializeTaskV10{})
	RegisterEvent(InitializeTaskV20{})
	RegisterEvent(InitializeTaskV30{})
	RegisterEvent(StartTaskV10{})
	RegisterEvent(StartTaskV20{})
	RegisterEvent(StartTaskV30{})
	RegisterEvent(StartTaskV40{})
	RegisterEvent(LogV10{})
	RegisterEvent(LogV20{})
	RegisterEvent(LogV30{})
	RegisterEvent(LogV40{})
	RegisterEvent(FinishGetV10{})
	RegisterEvent(FinishGetV20{})
	RegisterEvent(FinishGetV30{})
	RegisterEvent(FinishPutV10{})
	RegisterEvent(FinishPutV20{})
	RegisterEvent(FinishPutV30{})
	RegisterEvent(InitializeGetV10{})
	RegisterEvent(InitializePutV10{})
	RegisterEvent(FinishGetV40{})
	RegisterEvent(FinishPutV40{})
}

type Message struct {
	Event types.Event
}

type Envelope struct {
	Data    *json.RawMessage   `json:"data"`
	Event   types.EventType    `json:"event"`
	Version types.EventVersion `json:"version"`
}

func (m Message) MarshalJSON() ([]byte, error) {
	var envelope Envelope

	payload, err := json.Marshal(m.Event)
	if err != nil {
		return nil, err
	}

	envelope.Data = (*json.RawMessage)(&payload)
	envelope.Event = m.Event.EventType()
	envelope.Version = m.Event.Version()

	return json.Marshal(envelope)
}

func (m *Message) UnmarshalJSON(bytes []byte) error {
	var envelope Envelope

	err := json.Unmarshal(bytes, &envelope)
	if err != nil {
		return err
	}

	event, err := ParseEvent(envelope.Version, envelope.Event, *envelope.Data)
	if err != nil {
		return err
	}

	m.Event = event

	return nil
}

type UnknownEventTypeError struct {
	Type types.EventType
}

func (err UnknownEventTypeError) Error() string {
	return fmt.Sprintf("unknown event type: %s", err.Type)
}

type UnknownEventVersionError struct {
	Type          types.EventType
	Version       types.EventVersion
	KnownVersions []string
}

func (err UnknownEventVersionError) Error() string {
	return fmt.Sprintf(
		"unknown event version: %s version %s (supported versions: %s)",
		err.Type,
		err.Version,
		strings.Join(err.KnownVersions, ", "),
	)
}

func ParseEvent(version types.EventVersion, typ types.EventType, payload []byte) (types.Event, error) {
	versions, found := events[typ]
	if !found {
		return nil, UnknownEventTypeError{typ}
	}

	knownVersions := []string{}
	for v, parser := range versions {
		knownVersions = append(knownVersions, string(v))

		if v.IsCompatibleWith(version) {
			return parser(payload)
		}
	}

	return nil, UnknownEventVersionError{
		Type:          typ,
		Version:       version,
		KnownVersions: knownVersions,
	}
}
