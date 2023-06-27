package pubsub

import "sync"

type Broker[T any] struct {
	mutex       sync.Mutex
	subscribers map[string][]chan Message[T]
}

func NewBroker[T any]() *Broker[T] {
	return &Broker[T]{}
}

func (b *Broker[T]) Subscribe(queue string, subscriber chan Message[T]) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.subscribers == nil {
		b.subscribers = make(map[string][]chan Message[T])
	}

	b.subscribers[queue] = append(b.subscribers[queue], subscriber)
}

func (b *Broker[T]) Publish(queue string, data T) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if subscribers, found := b.subscribers[queue]; found {
		for _, subscriber := range subscribers {
			go func(subscriber chan Message[T]) {
				subscriber <- Message[T]{Queue: queue, Data: data}
			}(subscriber)
		}
	}
}
