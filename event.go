package goevents

// An `Event` represents a Event to be sent to a `Listener`. 
//
// Its receives a name for identifying the Event and 
// a callback function that actually executes the
// asynchronous code.
type Event struct {
    EventName   string
    Callback    func(data ...any)
}

func NewEvent(EventName string, Callback func(data ...any)) Event {
    return Event {
	EventName,
	Callback,
    }
}
