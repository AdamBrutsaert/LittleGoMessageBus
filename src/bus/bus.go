package bus

import "reflect"

type callable interface {
	Call(interface{})
}

type Function[T any] func(T)

func (fn Function[T]) Call(value interface{}) {
	fn(value.(T))
}

type messageBus struct {
	Subscribers map[reflect.Type][]callable
}

func (bus *messageBus) Subscribe(t reflect.Type, fn callable) {
	bus.Subscribers[t] = append(bus.Subscribers[t], fn)
}

var bus messageBus

func init() {
	bus.Subscribers = make(map[reflect.Type][]callable)
}

func getType[T any]() reflect.Type {
	var temp T
	return reflect.TypeOf(temp)
}

func Subscribe[T any](fn callable) {
	bus.Subscribers[getType[T]()] = append(bus.Subscribers[getType[T]()], fn)
}

func Send[T any](value T) {
	for _, fn := range bus.Subscribers[reflect.TypeOf(value)] {
		fn.Call(value)
	}
}
