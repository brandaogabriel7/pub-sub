package storage

import (
	"fmt"
	"sync"
	"time"

	"github.com/brandaogabriel7/pubsub/messages"
)

// FileMessageStorage is an implementation of IMessageStorage that stores messages in files.
type FileMessageStorage[T any] struct {
	filesMutex map[string]*sync.Mutex
	fileWriter FileWriter
}

// StoreMessage stores a message in a file.
func (f *FileMessageStorage[T]) StoreMessage(message messages.Message[T]) {
	func(message messages.Message[T]) {
		if fileMutex, found := f.filesMutex[message.Queue]; found {
			fileMutex.Lock()
			defer fileMutex.Unlock()
		} else {
			f.filesMutex[message.Queue] = &sync.Mutex{}
			f.filesMutex[message.Queue].Lock()
			defer f.filesMutex[message.Queue].Unlock()
		}
		filename := message.Queue + ".txt"
		timestamp := time.Now().Format(time.RFC3339)
		entry := fmt.Sprintf("[%s] %v\n", timestamp, message.Data)

		err := f.fileWriter.WriteToFile(filename, entry)

		if err != nil {
			fmt.Printf("error storing %v to file %v\n", message, err)
		}
	}(message)
}

// NewFileMessageStorage creates a new FileMessageStorage instance.
func NewFileMessageStorage[T any](fileWriter FileWriter) *FileMessageStorage[T] {
	return &FileMessageStorage[T]{filesMutex: make(map[string]*sync.Mutex), fileWriter: fileWriter}
}
