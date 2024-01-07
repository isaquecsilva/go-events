# go-events
go-events is a simple Go library for running asynchronous go code based on event-driven programming.

## Installation:
```
go get github.com/isaquecsilva/go-events
```

## Usage:

Example using a event **Listener** with no capacity limits:

```go
package main

import (
    "io"
    "os"
    "net/http"

    "github.com/isaquecsilva/go-events"
)

func main() {
    listener := goevents.NewListener()

    if err := listener.Add(
	    goevents.NewEvent("download", func(data any) {
            response, ok := data.(*http.Response)

            if ok {
                defer response.Body.Close()
                io.Copy(os.Stdout, response.Body)
            } else {
                panic("couldn't convert to http.Response object")
            }

        }),
    ); err != nil {
	    panic(err)
    }
   
    goevents.Emit(listener, "download", func() *http.Response {
	    response, err := http.Get("https://go.dev/")
        if err != nil {
            panic(err)
        }

        return response
    }())

    listener.Wait()
}
```

Example using a event **Listener** with specific capacity:

```go
package main

import (
	"fmt"

	goevents "github.com/isaquecsilva/go-events"
)

func main() {
    // Listener cannot have more than 2 events in its
    // buffer.
	listener, err := goevents.NewListenerWithCapacity(2)
	if err != nil {
		panic(err)
	}

	event := goevents.NewEvent("simple-event", func(data any) {
		fmt.Printf("parameter is %s\n", data.(string))
	})

	err = listener.Add(event)
	if err != nil {
		panic(err)
	}

	event = goevents.NewEvent(
		"even-odd",
		func(data any) {
			number := data.(int)

			switch number % 2 {
			case 0:
				println("An even number received")
			default:
				println("An odd number received")
			}
		},
	)

	err = listener.Add(event)
	if err != nil {
		panic(err)
	}

	goevents.Emit(listener, "simple-event", "simple")
	goevents.Emit(listener, "even-odd", 25)
	listener.Wait()
}
```