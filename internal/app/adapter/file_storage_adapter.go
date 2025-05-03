package adapter

type FileStorageAdapter interface {
	StoreFile(fileName, fileType string, fileContent []byte) (string, error)
}
