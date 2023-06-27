package storage

// FileWriter is an interface that defines the methods that a file writer implementation must implement.
type FileWriter interface {
	// WriteToFile writes an entry to a file.
	WriteToFile(filename string, entry string) error
}
