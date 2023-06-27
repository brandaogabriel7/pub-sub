package pubsub_test

import (
	"fmt"
	"testing"

	"github.com/brandaogabriel7/pubsub"
	"github.com/brandaogabriel7/pubsub/messages"
)

type messageStorageMock[T comparable] struct {
	messagesStored []messages.Message[T]
}

func (m *messageStorageMock[T]) StoreMessage(message messages.Message[T]) {
	if m.messagesStored == nil {
		m.messagesStored = []messages.Message[T]{}
	}
	m.messagesStored = append(m.messagesStored, message)
}

func (m *messageStorageMock[T]) HasStoredMessage(message messages.Message[T]) bool {
	fmt.Printf("checking if message %v is stored in %v\n", message, m.messagesStored)
	for _, storedMessage := range m.messagesStored {
		if storedMessage.Queue == message.Queue && storedMessage.Data == message.Data {
			return true
		}
	}
	return false
}

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
			broker := pubsub.NewBroker[string](nil)

			subscribers := []chan messages.Message[string]{
				make(chan messages.Message[string]),
				make(chan messages.Message[string]),
				make(chan messages.Message[string]),
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
			broker := pubsub.NewBroker[int](nil)

			subscribers := []chan messages.Message[int]{
				make(chan messages.Message[int]),
				make(chan messages.Message[int]),
				make(chan messages.Message[int]),
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

func TestStructBroker(t *testing.T) {
	var testCases = []struct {
		queue string
		data  struct {
			Name string
			Age  int
		}
	}{
		{"topic", struct {
			Name string
			Age  int
		}{"John Doe", 30}},
		{"topic2", struct {
			Name string
			Age  int
		}{"Jane Doe", 25}},
		{"another topic", struct {
			Name string
			Age  int
		}{"John Smith", 40}},
	}

	for _, testCase := range testCases {
		testname := fmt.Sprintf("publish to %s queue", testCase.queue)
		t.Run(testname, func(t *testing.T) {
			broker := pubsub.NewBroker[struct {
				Name string
				Age  int
			}](nil)

			subscribers := []chan messages.Message[struct {
				Name string
				Age  int
			}]{
				make(chan messages.Message[struct {
					Name string
					Age  int
				}]),
				make(chan messages.Message[struct {
					Name string
					Age  int
				}]),
				make(chan messages.Message[struct {
					Name string
					Age  int
				}]),
			}

			for _, subscriber := range subscribers {
				broker.Subscribe(testCase.queue, subscriber)
			}

			broker.Publish(testCase.queue, testCase.data)

			for _, subscriber := range subscribers {
				if message := <-subscriber; message.Queue != testCase.queue || message.Data != testCase.data {
					t.Errorf("expected message to be {Queue: \"%s\", Data: %v}, got %v", testCase.queue, testCase.data, message)
				}
			}
		})
	}
}

func TestMultipleMessagesPublishing(t *testing.T) {
	var testCases = []struct {
		queue string
		data  []string
	}{
		{"topic", []string{"hello world", "hello world 2", "hello world 3"}},
		{"topic2", []string{"hello world 4", "hello world 5", "hello world 6"}},
		{"another topic", []string{"hello world 7", "hello world 8", "hello world 9"}},
	}

	for _, testCase := range testCases {
		testname := fmt.Sprintf("publish to %s queue", testCase.queue)
		t.Run(testname, func(t *testing.T) {
			broker := pubsub.NewBroker[string](nil)

			subscribers := []chan messages.Message[string]{
				make(chan messages.Message[string]),
				make(chan messages.Message[string]),
				make(chan messages.Message[string]),
			}

			for _, subscriber := range subscribers {
				broker.Subscribe(testCase.queue, subscriber)
			}

			for _, data := range testCase.data {
				broker.Publish(testCase.queue, data)

				for _, subscriber := range subscribers {
					if message := <-subscriber; message.Queue != testCase.queue || message.Data != data {
						t.Errorf("expected message for to be {Queue: \"%s\", Data: \"%s\"}, got %v", testCase.queue, data, message)
					}
				}
			}
		})
	}
}

func TestBrokerMessageStorage(t *testing.T) {
	var testCases = []struct {
		queue string
		data  string
	}{
		{"topic", "hello world"},
		// {"topic2", "hello world 2"},
		// {"another topic", "hello world 3"},
	}

	for _, testCase := range testCases {
		testname := fmt.Sprintf("publish to %s queue", testCase.queue)
		t.Run(testname, func(t *testing.T) {
			messageStorage := &messageStorageMock[string]{}
			broker := pubsub.NewBroker[string](messageStorage)

			subscribers := []chan messages.Message[string]{
				make(chan messages.Message[string]),
				make(chan messages.Message[string]),
				make(chan messages.Message[string]),
			}

			for _, subscriber := range subscribers {
				broker.Subscribe(testCase.queue, subscriber)
			}

			broker.Publish(testCase.queue, testCase.data)

			if !messageStorage.HasStoredMessage(messages.Message[string]{Queue: testCase.queue, Data: testCase.data}) {
				t.Errorf("expected message %v to be stored", testCase)
			}
		})
	}
}
