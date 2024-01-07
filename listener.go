package goevents

import (
    "fmt"
    "sync"
)

const (
    not_found_event = -1
)

var (
    // mutex for when removing event from listener buffer
    mu sync.Mutex
)

// Listener represents the listener object that will contain events and execute event callbacks when they are called. 
type Listener struct {
    events    []Event
    capacity  uint64
    wg        sync.WaitGroup
}

// NewListener creates a new Listener object and returns a pointer to it.
func NewListener() *Listener {
    return &Listener{}
}

// NewListenerWithCapacity tries to create a new event listener with certain storage capacity for the amount of events that will be added.
// It returns an error if capacity is 0.
func NewListenerWithCapacity(capacity uint64) (*Listener, error) {
    if capacity == 0 {
	return nil, fmt.Errorf("Listener capacity cannot be %d", capacity)
    }

    return &Listener {
	capacity: capacity,
	events: make([]Event, 0, capacity),
    }, nil
}

// GetCap returns the Listener capacity. If capacity was not set, as would be with NewListenerWithCapacity, the return value is 0.
func (l *Listener) GetCap() int {
    return cap(l.events)
}

// GetLen returns Listener current length of added events.
func (l *Listener) GetLen() int {
    return len(l.events)
}

// Add adds an Event to Listener. Whether the event is already set on Listener, or the addition exceeds Listener capacity (when it's set to have one), Add returns a non-nil error value. Otherwise, the error will be nil.
func (l *Listener) Add(event Event) error {
    if len(l.events) == cap(l.events) && l.capacity > 0 {
	return fmt.Errorf("event appending exceeds Listener buffer capacity")
    } else {
	
	if event, _ := l.find(event.EventName); event != nil {
	    return fmt.Errorf("event already set")
	}

	l.events = append(l.events, event)
    }

    return nil
}

// Remove removes an Event that resides inside the Listener. If the event is not found in Listener, it returns an non-nil error value. 
func (l *Listener) Remove(eventName string) error {

    if _, eventIndex := l.find(eventName); eventIndex != not_found_event {

	switch eventIndex {
	case 0:
	    mu.Lock()
	    l.events = l.events[1:]
	    mu.Unlock()
	case len(l.events) - 1:
	    mu.Lock()
	    l.events = l.events[:eventIndex]
	    mu.Unlock()
	default:
	    mu.Lock()
	    l.events = append(l.events, l.events[:eventIndex]...)
	    l.events = append(l.events, l.events[eventIndex+1:]...)
	    mu.Unlock()
	}
    } else {
	return fmt.Errorf("specified event not found in listener buffer")
    }

    return nil
}

// Wait waits for an emitted Event to be run before continues.
func (l *Listener) Wait() {
    l.wg.Wait()
}

func (l *Listener) find(eventName string) (*Event, int) {
    for index, event := range l.events {
	if eventName == event.EventName {
	    return &event, index
	}
    }

    return nil, -1
}

func (l *Listener) receive(eventName string, data any) error {
    // tries to found the event inside listener events list
    if event, _ := l.find(eventName); event != nil {
	l.wg.Add(1)

	go func() {
	    (*event).Callback(data)
	    l.wg.Done()
	}()

	return nil
    }

    return fmt.Errorf("event not found in listener")
}
