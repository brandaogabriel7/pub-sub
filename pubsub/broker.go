package pubsub

import (
	"sync"

	"github.com/brandaogabriel7/pubsub/messages"
	"github.com/brandaogabriel7/pubsub/storage"
)

type Broker[T comparable] struct {
	mutex          sync.Mutex
	subscribers    map[string][]chan messages.Message[T]
	messageStorage storage.IMessageStorage[T]
}

func NewBroker[T comparable](messageStorage storage.IMessageStorage[T]) *Broker[T] {
	if messageStorage != nil {
		return &Broker[T]{mutex: sync.Mutex{}, messageStorage: messageStorage}
	}
	return &Broker[T]{mutex: sync.Mutex{}}
}

func (b *Broker[T]) Subscribe(queue string, subscriber chan messages.Message[T]) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.subscribers == nil {
		b.subscribers = make(map[string][]chan messages.Message[T])
	}

	b.subscribers[queue] = append(b.subscribers[queue], subscriber)
}

func (b *Broker[T]) Publish(queue string, data T) {
	message := messages.Message[T]{Queue: queue, Data: data}
	if b.messageStorage != nil {
		b.messageStorage.StoreMessage(message)
	}

	b.mutex.Lock()
	defer b.mutex.Unlock()

	if subscribers, found := b.subscribers[queue]; found {
		for _, subscriber := range subscribers {
			go func(subscriber chan messages.Message[T]) {
				subscriber <- message
			}(subscriber)
		}
	}
}
