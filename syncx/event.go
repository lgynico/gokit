package syncx

type Event[T any] struct {
	eventC chan T
}

func NewEvent[T any](buf int) *Event[T] {
	return &Event[T]{
		eventC: make(chan T, buf),
	}
}

func (e *Event[T]) Wait() T {
	return <-e.eventC
}

func (e *Event[T]) Notify(t T) {
	e.eventC <- t
}
