package pubsub

import (
	"sync"

	"github.com/brandaogabriel7/pubsub/messages"
	"github.com/brandaogabriel7/pubsub/storage"
)

// Broker is a thread-safe publisher/subscriber implementation. It allows subscribers to subscribe to a queue and receive messages
// published to that queue and also allows publishers to publish messages to a queue.
// Broker can be initialized with a message storage implementation. If a message storage is provided, all messages
// published to a queue will be stored in the storage.
type Broker[T comparable] struct {
	mutex          sync.Mutex
	subscribers    map[string][]chan messages.Message[T]
	messageStorage storage.IMessageStorage[T]
}

// NewBroker creates a new Broker instance.
func NewBroker[T comparable](messageStorage storage.IMessageStorage[T]) *Broker[T] {
	if messageStorage != nil {
		return &Broker[T]{mutex: sync.Mutex{}, messageStorage: messageStorage}
	}
	return &Broker[T]{mutex: sync.Mutex{}}
}

// Subscribe subscribes a channel to a queue. All messages published to the queue will be sent to the channel.
func (b *Broker[T]) Subscribe(queue string, subscriber chan messages.Message[T]) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.subscribers == nil {
		b.subscribers = make(map[string][]chan messages.Message[T])
	}

	b.subscribers[queue] = append(b.subscribers[queue], subscriber)
}

// Publish publishes a message to a queue. All subscribers to the queue will receive the message.
// If a message storage was provided, the message will be stored in the storage.
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
