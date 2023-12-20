package fs

type FileSystem interface {
	Read(filePath string) ([]byte, error)
	Write(filePath string, base64Content string) error
	Remove(filePath string) error
}
