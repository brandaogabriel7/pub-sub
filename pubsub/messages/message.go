package messages

type Message[T comparable] struct {
	Queue string
	Data  T
}
