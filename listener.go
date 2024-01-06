package goevents

import (
    "fmt"
    "slices"
    "sync"
)

const (
    not_found_event = -1
)

var (
    // mutex for when removing event from listener buffer
    mu sync.Mutex
)

// `Listener` represents the Listener for Events.
type Listener struct {
    events    []Event
}

func NewListener() *Listener {
    return &Listener{}
}

func NewListenerWithCapacity(cap int) (*Listener, error) {
    if cap <= 0 {
	return nil, fmt.Errorf("Listener capacity cannot be %d", cap)
    }

    return &Listener {
	events: make([]Event, 0, cap),
    }
}

func (l *Listener) Add(event Event) error {
    if slices.Index(l.events, event) != not_found_event {
	return fmt.Errorf("event already set")
    }

    if len(l.events) == cap(l.events) {
	return fmt.Errorf("event appending exceeds Listener buffer capacity")
    } else {
	l.events = append(l.events, event)
    }

    return nil
}

func (l *Listener) Remove(event Event) error {
    if eventIndex := slices.Index(l.events, event); eventIndex != not_found_event {
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
	    l.events = append(l.events, l.events[:eventIndex]..., l.events[eventIndex+1:]...)
	    mu.Unlock()
	}
    } else {
	return fmt.Errorf("specified event not found in listener buffer")
    }

    return nil
}

func (l *Listener) receive(eventName string, data any) error {
    // try to found the event inside listener events list
    for index, _ := range l.events {

	if l.events[index].Name == eventName {
	    // Whether event is found, then we call event call method
	    // on a goroutine
	    go l.events[index](data)
	    return nil
	}
    }

    return fmt.Errorf("event not found in listener")
}
