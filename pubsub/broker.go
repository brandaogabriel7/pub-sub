package pubsub

type Broker[T any] struct {
	Subscribers map[string][]chan Message[T]
}

func NewBroker[T any]() *Broker[T] {
	return &Broker[T]{}
}

func (b *Broker[T]) Subscribe(queue string, subscriber chan Message[T]) {
	if b.Subscribers == nil {
		b.Subscribers = make(map[string][]chan Message[T])
	}

	b.Subscribers[queue] = append(b.Subscribers[queue], subscriber)
}
