package messages

// Message is a generic struct for messages to be published to a queue.
type Message[T comparable] struct {
	Queue string
	Data  T
}
