package storage_test

import (
	"os"
	"testing"

	"github.com/brandaogabriel7/pubsub/storage"
)

func TestOsFileWriter(t *testing.T) {
	var testCases = []struct {
		filename string
		entry    string
	}{
		{"test1.txt", "test1"},
		{"test2.txt", "test2"},
		{"test3.txt", "test3"},
	}

	for _, testCase := range testCases {
		osFileWriter := storage.NewOsFileWriter()

		err := osFileWriter.WriteToFile(testCase.filename, testCase.entry)
		if err != nil {
			t.Error(err)
		}

		os.Remove(testCase.filename)
	}
}
