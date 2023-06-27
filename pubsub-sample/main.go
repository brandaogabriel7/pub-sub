package main

import (
	"fmt"
	"time"

	"github.com/brandaogabriel7/pubsub"
	"github.com/brandaogabriel7/pubsub/messages"
	"github.com/brandaogabriel7/pubsub/storage"
)

func main() {
	fileWriter := storage.NewOsFileWriter()
	fileMessageStorage := storage.NewFileMessageStorage[string](fileWriter)
	broker := pubsub.NewBroker[string](fileMessageStorage)

	subscriber1 := make(chan messages.Message[string])
	defer close(subscriber1)
	subscriber2 := make(chan messages.Message[string])
	defer close(subscriber2)

	broker.Subscribe("queue1", subscriber1)
	broker.Subscribe("queue1", subscriber2)
	broker.Subscribe("queue2", subscriber1)

	go func() {
		for msg := range subscriber1 {
			fmt.Printf("Subscriber 1: %v\n", msg)
		}
	}()

	go func() {
		for msg := range subscriber2 {
			fmt.Printf("Subscriber 2: %v\n", msg)
		}
	}()

	broker.Publish("queue1", "Hello world!")
	broker.Publish("queue2", "Hi world!")
	broker.Publish("queue2", "bão demais uai")

	time.Sleep(3 * time.Second)
}
