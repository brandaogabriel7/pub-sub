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

func (b *Broker[T]) Publish(queue string, data T) {
	if b.Subscribers == nil {
		return
	}

	go func() {
		for _, subscriber := range b.Subscribers[queue] {
			subscriber <- Message[T]{Queue: queue, Data: data}
		}
	}()
}
