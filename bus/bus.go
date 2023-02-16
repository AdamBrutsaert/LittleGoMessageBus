package bus

import "reflect"

var bus messageBus

type messageBus struct {
	Subscribers map[reflect.Type][]callable
}

func init() {
	bus.Subscribers = make(map[reflect.Type][]callable)
}

func Subscribe[T any](fn func(T)) {
	var temp T
	t := reflect.TypeOf(temp)

	var proxy function[T] = fn
	bus.Subscribers[t] = append(bus.Subscribers[t], proxy)
}

func Send[T any](value T) {
	for _, fn := range bus.Subscribers[reflect.TypeOf(value)] {
		fn.Call(value)
	}
}
