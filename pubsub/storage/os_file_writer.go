package storage

import "os"

// OsFileWriter is an implementation of IFileWriter that writes to a file in the OS.
type OsFileWriter struct{}

// NewOsFileWriter creates a new OsFileWriter instance.
func NewOsFileWriter() *OsFileWriter {
	return &OsFileWriter{}
}

// WriteToFile writes an entry to a file.
func (w *OsFileWriter) WriteToFile(filename string, entry string) error {

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(entry); err != nil {
		return err
	}

	return nil
}
