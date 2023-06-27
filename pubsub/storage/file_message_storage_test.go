package storage_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/brandaogabriel7/pubsub/messages"
	"github.com/brandaogabriel7/pubsub/storage"
)

type fileWriterMock struct {
	entries []string
}

func (f *fileWriterMock) WriteToFile(filename string, entry string) error {
	f.entries = append(f.entries, entry)
	return nil
}

func (f *fileWriterMock) HasStoredMessage(message string) bool {
	for _, entry := range f.entries {
		if strings.Contains(entry, message) {
			return true
		}
	}
	return false
}

func TestStoreMessageInFile(t *testing.T) {
	var testCases = []struct {
		queue    string
		messages []messages.Message[string]
	}{
		{"topic", []messages.Message[string]{
			{Queue: "topic", Data: "hello world"},
			{Queue: "topic", Data: "hello world 2"},
			{Queue: "topic", Data: "hello world 3"},
		}},
		{"topic2", []messages.Message[string]{
			{Queue: "topic2", Data: "hello world"},
			{Queue: "topic2", Data: "hello world 2"},
			{Queue: "topic2", Data: "hello world 3"},
		}},
		{"another topic", []messages.Message[string]{
			{Queue: "another topic", Data: "hello world"},
			{Queue: "another topic", Data: "hello world 2"},
			{Queue: "another topic", Data: "hello world 3"},
		}},
	}

	for _, testCase := range testCases {
		testname := fmt.Sprintf("store messages in %s queue", testCase.queue)
		t.Run(testname, func(t *testing.T) {
			fileWriterMock := &fileWriterMock{}
			fileStorage := storage.NewFileMessageStorage[string](fileWriterMock)
			for _, message := range testCase.messages {
				fileStorage.StoreMessage(message)
			}

			for _, message := range testCase.messages {
				if !fileWriterMock.HasStoredMessage(message.Data) {
					t.Errorf("message %v was not stored in file", message)
				}
			}
		})
	}
}
