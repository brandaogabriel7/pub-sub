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
