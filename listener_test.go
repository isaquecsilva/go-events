package goevents

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestNewListener(test *testing.T) {
    test.Run("should get a valid Listener object", func(t *testing.T) {
	l := NewListener()
	assert.Equal(t, &Listener{}, l, "Listener obects are not equal")
    })
}

func TestNewListenerWithCapacity(test *testing.T) {
    test.Run("should get a valida Listener with buffering capacity", func(t *testing.T) {
	_, err := NewListenerWithCapacity(2)
	if err != nil {
	    t.FailNow()
	}
    })


    test.Run("should get an error when creating Listener with cap", func(t *testing.T) {
	_, err := NewListenerWithCapacity(0)
	if err == nil {
	    t.FailNow()
	}
    })

}

func TestGetCap(test *testing.T) {
    test.Run("expects a valid number of listener buffering capacity", func(t *testing.T) {
	l, _ := NewListenerWithCapacity(10)
	
	assert.Equal(t, 10, l.GetCap())
    })
}

func TestGetLen(test *testing.T) {
    test.Run("expects a valid number of listener events length", func(t *testing.T) {
	l := NewListener()
	l.Add(NewEvent("event", func(data any){}))

	assert.Equal(t, 1, l.GetLen())
    })
}

func TestAdd(test *testing.T) {
    var event = NewEvent("sample-event", func(data any) {})
    
    test.Run("should get a <nil> error", func(t *testing.T) {
	l, _ := NewListenerWithCapacity(1)	
	assert.Nil(t, l.Add(event))
    })

    test.Run("should get non-nil error by exceeding listener cap", func(t *testing.T) {
	l, _ := NewListenerWithCapacity(1)
	l.Add(event)
	err := l.Add(NewEvent("sample-event-2", func(data any) {}))
	assert.NotNil(t, err)
    })

    test.Run("expects a non-nil error cause will try add a already set event", func(t *testing.T) {
	l := NewListener()
	l.Add(event)
	assert.NotNil(t, l.Add(event))
    })
}

func TestRemove(test *testing.T) {
    var event Event = NewEvent("sample-event", func(data any) {})

    test.Run("should remove a added event", func(t *testing.T) {
	l := NewListener()
	l.Add(event)

	err := l.Remove(event.EventName)

	assert.Nil(t, err)
    })

    test.Run("should fail on removing an event that is not found in listener", func(t *testing.T) {
	l := NewListener()
	assert.NotNil(t, l.Remove(event.EventName))
    })
}

func TestFind(test *testing.T) {
    var event = NewEvent("sample-event", func(data any) {})

    test.Run("expects find a event inside listener", func(t *testing.T) {
	l := NewListener()
	l.Add(event)
	_, index := l.find(event.EventName)
	
	assert.NotEqual(t, index, not_found_event)
    })

    test.Run("expects a not_found_event index", func(t *testing.T) {
	l := NewListener()
	l.Add(event)
	_, index := l.find("non-existing-event")

	assert.Equal(t, index, not_found_event)
    })
    
}

func TestReceive(test *testing.T) {
    var event = NewEvent("sample-event", func(data any) {})

    test.Run("should get a nil error", func(t *testing.T) {
	l := NewListener()
	l.Add(event)
	err := l.receive("sample-event", nil)

	assert.Nil(t, err)
    })

    test.Run("expects a non-nil error return value", func(t *testing.T) {
	l := NewListener()
	assert.NotNil(t, l.receive("non-existing-event", nil))
    })
}
