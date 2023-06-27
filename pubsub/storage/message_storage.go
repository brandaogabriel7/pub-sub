package storage

import "github.com/brandaogabriel7/pubsub/messages"

type IMessageStorage[T comparable] interface {
	StoreMessage(message messages.Message[T])
}
