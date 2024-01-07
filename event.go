package goevents

// An Event represents a event to be sent to a Listener.
//
// It receives a name for identifying the Event and a callback function that actually executes the asynchronous code.
type Event struct {
    EventName   string
    Callback    func(data any)
}

// NewEvent creates a new Event object. Accepting a event's name and a callback function.
func NewEvent(EventName string, Callback func(data any)) Event {
    return Event {
	EventName,
	Callback,
    }
}
