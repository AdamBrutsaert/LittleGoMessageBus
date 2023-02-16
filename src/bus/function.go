package bus

type callable interface {
	Call(interface{})
}

type function[T any] func(T)

func (fn function[T]) Call(value interface{}) {
	fn(value.(T))
}
