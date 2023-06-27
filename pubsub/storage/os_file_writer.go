package storage

import "os"

type OsFileWriter struct{}

func NewOsFileWriter() *OsFileWriter {
	return &OsFileWriter{}
}

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
