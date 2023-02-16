package main

import (
	"fmt"

	"github.com/AdamBrutsaert/LittleGoMessageBus/bus"
)

func main() {
	// Subscribe to a type
	bus.Subscribe(func(value int) {
		fmt.Println("Value is:", value)
	})

	// Send messages
	bus.Send("") // No function handles it
	bus.Send(3)  // A function handles it!
}
