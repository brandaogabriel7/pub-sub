package pubsub_test

import (
	"fmt"
	"testing"

	"github.com/brandaogabriel7/pubsub"
)

func TestStringBroker(t *testing.T) {
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

			subscribers := []chan pubsub.Message[string]{
				make(chan pubsub.Message[string]),
				make(chan pubsub.Message[string]),
				make(chan pubsub.Message[string]),
			}

			for _, subscriber := range subscribers {
				broker.Subscribe(testCase.queue, subscriber)
			}

			broker.Publish(testCase.queue, testCase.data)

			for _, subscriber := range subscribers {
				if message := <-subscriber; message.Queue != testCase.queue || message.Data != testCase.data {
					t.Errorf("expected message to be {Queue: \"%s\", Data: \"%s\"}, got %v", testCase.queue, testCase.data, message)
				}
			}
		})
	}
}

func TestIntBroker(t *testing.T) {
	var testCases = []struct {
		queue string
		data  int
	}{
		{"topic", 1},
		{"topic2", 2},
		{"another topic", 3},
	}

	for _, testCase := range testCases {
		testname := fmt.Sprintf("publish to %s queue", testCase.queue)
		t.Run(testname, func(t *testing.T) {
			broker := pubsub.NewBroker[int]()

			subscribers := []chan pubsub.Message[int]{
				make(chan pubsub.Message[int]),
				make(chan pubsub.Message[int]),
				make(chan pubsub.Message[int]),
			}

			for _, subscriber := range subscribers {
				broker.Subscribe(testCase.queue, subscriber)
			}

			broker.Publish(testCase.queue, testCase.data)

			for _, subscriber := range subscribers {
				if message := <-subscriber; message.Queue != testCase.queue || message.Data != testCase.data {
					t.Errorf("expected message to be {Queue: \"%s\", Data: %d}, got %v", testCase.queue, testCase.data, message)
				}
			}
		})
	}
}
