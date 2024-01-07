package goevents

import (
    "testing"
    "reflect"
)

func TestNewEvent(test *testing.T) {
    test.Run("Expects a valid and equal Event object", func(t *testing.T) {
	var event Event = NewEvent("someevent", nil)

	if reflect.DeepEqual(Event{
	    EventName: "someevent",
	    Callback: nil,
	}, event) != true {
	    t.FailNow()
	}
    })
}
