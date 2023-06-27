package pubsub_test

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/brandaogabriel7/pubsub"
	"github.com/brandaogabriel7/pubsub/messages"
	"github.com/brandaogabriel7/pubsub/storage"
)

func TestPubSubWithOsFileStorage(t *testing.T) {
	queue := "integration_test_queue"
	fileWriter := storage.NewOsFileWriter()
	fileMessageStorage := storage.NewFileMessageStorage[string](fileWriter)
	broker := pubsub.NewBroker[string](fileMessageStorage)

	subscriber := make(chan messages.Message[string])
	defer close(subscriber)

	go func() {
		for message := range subscriber {
			t.Logf("Received message: %s", message)
		}
	}()

	broker.Subscribe(queue, subscriber)

	messages := []string{"Hello world!", "Hi world!", "bão demais uai"}

	for _, message := range messages {
		broker.Publish(queue, message)
	}

	// wait for messages to be written to file
	time.Sleep(1 * time.Second)

	data, err := os.ReadFile(queue + ".txt")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	lines := strings.Split(string(data), "\n")
	for i, message := range messages {
		if !strings.Contains(lines[i], message) {
			t.Errorf("Expected line to contain message %s, but got %s", message, lines[i])
		}
	}

	// delete test file
	os.Remove(queue + ".txt")
}
