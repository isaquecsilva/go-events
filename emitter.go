package goevents

import "time"

// `Emit` dispatches an event to the specified listener.
// It returns a non nil error if fails.
func Emit(listener *Listener, eventName string, data any) error {
    var err error = listener.receive(eventName, data)

    // Adding some delay for when emitting events 
    // in certain call order. Not using this, can make
    // emitting events chain (one emission followed by another one)
    // may gets its corresponding events in listener to be run
    // without the call order.
    time.Sleep(time.Nanosecond)

    return err
}
