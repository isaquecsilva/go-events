package goevents

import (
    "testing"
)

func TestEmit(test *testing.T) {
    test.Run("Should get no error on event emitting", func(t *testing.T) {
	listener := NewListener()
	listener.Add(NewEvent("test-call", func(data any) {}))

	if err := Emit(listener, "test-call", nil); err != nil {
	    t.FailNow()
	}

    })
}
