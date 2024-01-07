package main


import (
    "fmt"
    
    "github.com/isaquecsilva/go-events"
)

func main() {
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
