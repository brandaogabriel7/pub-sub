package storage

import "github.com/brandaogabriel7/pubsub/messages"

// IMessageStorage is an interface that defines the methods that a message storage implementation must implement.
type IMessageStorage[T any] interface {
	StoreMessage(message messages.Message[T])
}
