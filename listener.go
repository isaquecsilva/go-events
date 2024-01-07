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

// `Listener` represents the Listener for Events.
type Listener struct {
    events    []Event
    capacity  uint64
    wg        sync.WaitGroup
}

func NewListener() *Listener {
    return &Listener{}
}

func NewListenerWithCapacity(capacity uint64) (*Listener, error) {
    if capacity == 0 {
	return nil, fmt.Errorf("Listener capacity cannot be %d", capacity)
    }

    return &Listener {
	capacity: capacity,
	events: make([]Event, 0, capacity),
    }, nil
}

func (l *Listener) GetCap() int {
    return cap(l.events)
}

func (l *Listener) GetLen() int {
    return len(l.events)
}

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
    // try to found the event inside listener events list
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
