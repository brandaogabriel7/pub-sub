package storage

type FileWriter interface {
	WriteToFile(filename string, entry string) error
}
