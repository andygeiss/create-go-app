package event_test

import (
	"testing"

	"github.com/andygeiss/create-go-app/pkg/assert"
	"github.com/andygeiss/create-go-app/pkg/event"
)

func TestEventBusPublish(t *testing.T) {
	val := false
	bus := event.NewBus()
	bus.Subscribe("foo", func(data interface{}) {
		val = true
	})
	bus.Publish("foo", "bar")
	assert.That("publish should change value to true", t, val, true)
}

func TestEventBusSubscribe(t *testing.T) {
	bus := event.NewBus()
	bus.Subscribe("foo", func(data interface{}) {})
	assert.That("subscribe should add one function to topic foo", t, len(bus.TopicSubscribers["foo"]), 1)
}
