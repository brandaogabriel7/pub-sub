package pubsub

type Message[T any] struct {
	Queue string
	Data  T
}
