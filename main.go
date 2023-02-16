package main

import (
	"fmt"
	"reflect"
)

type Callable interface {
	Call(interface{})
}

type Function[T any] func(T)

func (fn Function[T]) Call(value interface{}) {
	fn(value.(T))
}

type MessageBus struct {
	Subscribers map[reflect.Type][]Callable
}

func (bus *MessageBus) Subscribe(t reflect.Type, fn Callable) {
	bus.Subscribers[t] = append(bus.Subscribers[t], fn)
}

var bus MessageBus

func init() {
	bus.Subscribers = make(map[reflect.Type][]Callable)
}

func getType[T any]() reflect.Type {
	var temp T
	return reflect.TypeOf(temp)
}

func Subscribe[T any](fn Callable) {
	bus.Subscribers[getType[T]()] = append(bus.Subscribers[getType[T]()], fn)
}

func Send[T any](value T) {
	for _, fn := range bus.Subscribers[reflect.TypeOf(value)] {
		fn.Call(value)
	}
}

func main() {
	// Create a handler for a certain type
	var fn Function[int] = func(value int) {
		fmt.Println("Value is: ", value)
	}

	// Subscribe to a type
	Subscribe[int](fn)

	// Send messages
	Send("") // No function handles it
	Send(3)  // A function handles it!
}
