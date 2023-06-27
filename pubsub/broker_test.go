package pubsub_test

import (
	"fmt"
	"testing"

	"github.com/brandaogabriel7/pubsub"
)

func TestStringBrokerSubscribe(t *testing.T) {
	var testCases = []struct {
		queue string
	}{
		{"topic"},
		{"topic2"},
		{"another topic"},
	}

	for _, testCase := range testCases {
		testname := fmt.Sprintf("subscribe to %s queue", testCase.queue)
		t.Run(testname, func(t *testing.T) {
			broker := pubsub.NewBroker[string]()

			subscriber := make(chan pubsub.Message[string])

			broker.Subscribe(testCase.queue, subscriber)

			if subscribers, found := broker.Subscribers[testCase.queue]; found {
				for _, sub := range subscribers {
					if sub == subscriber {
						return
					}
				}
			}

			t.Errorf("subscriber was not found in broker.Subscribers[\"%s\"]", testCase.queue)
		})
	}
}

func TestStringBrokerPublish(t *testing.T) {
	var testCases = []struct {
		queue string
		data  string
	}{
		{"topic", "hello world"},
		{"topic2", "hello world 2"},
		{"another topic", "hello world 3"},
	}

	for _, testCase := range testCases {
		testname := fmt.Sprintf("publish to %s queue", testCase.queue)
		t.Run(testname, func(t *testing.T) {
			broker := pubsub.NewBroker[string]()

			subscriber := make(chan pubsub.Message[string])

			broker.Subscribe(testCase.queue, subscriber)

			broker.Publish(testCase.queue, testCase.data)

			if message := <-subscriber; message.Queue != testCase.queue || message.Data != testCase.data {
				t.Errorf("expected message to be {Queue: \"%s\", Data: \"%s\"}, got %v", testCase.queue, testCase.data, message)
			}
		})
	}
}
