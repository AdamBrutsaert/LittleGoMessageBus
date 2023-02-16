package main

import (
	"fmt"

	"github.com/AdamBrutsaert/LittleGoMessageBus/bus"
)

func main() {
	// Create a handler for a certain type
	var fn bus.Function[int] = func(value int) {
		fmt.Println("Value is:", value)
	}

	// Subscribe to a type
	bus.Subscribe[int](fn)

	// Send messages
	bus.Send("") // No function handles it
	bus.Send(3)  // A function handles it!
}
