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

	return ( response )
    }())

    listener.Wait()
}
