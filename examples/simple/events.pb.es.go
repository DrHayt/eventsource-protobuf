package main

import (
	"fmt"
	"time"

	"github.com/altairsix/eventsource"
	"github.com/gogo/protobuf/proto"
)

type serializer struct {
}

func (s *serializer) MarshalEvent(event eventsource.Event) (eventsource.Record, error) {
	container := &EventContainer{}

	switch v := event.(type) {

	case *A:
		container.Type = 2
		container.Ma = v

	case *B:
		container.Type = 3
		container.Mb = v

	default:
		return eventsource.Record{}, fmt.Errorf("Unhandled type, %v", event)
	}

	data, err := proto.Marshal(container)
	if err != nil {
		return eventsource.Record{}, err
	}

	return eventsource.Record{
		Version: event.EventVersion(),
		Data:    data,
	}, nil
}

func (s *serializer) UnmarshalEvent(record eventsource.Record) (eventsource.Event, error) {
	container := &EventContainer{};
	err := proto.Unmarshal(record.Data, container)
	if err != nil {
		return nil, err
	}

	var event interface{}
	switch container.Type {

	case 2:
		event = container.Ma

	case 3:
		event = container.Mb

	default:
		return nil, fmt.Errorf("Unhandled type, %v", container.Type)
	}

	return event.(eventsource.Event), nil
}

func NewSerializer() eventsource.Serializer {
	return &serializer{}
}

func (m *A) AggregateID() string { return m.Id }
func (m *A) EventVersion() int   { return int(m.Version) }
func (m *A) EventAt() time.Time  { return time.Unix(m.At, 0) }

func (m *B) AggregateID() string { return m.Id }
func (m *B) EventVersion() int   { return int(m.Version) }
func (m *B) EventAt() time.Time  { return time.Unix(m.At, 0) }

