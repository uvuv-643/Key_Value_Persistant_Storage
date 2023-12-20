package fs

import (
	"encoding/base64"
	"os"
)

type LocalFileSystem struct {
	rootPath string
}

func NewLocalFileSystem(rootPath string) *LocalFileSystem {
	return &LocalFileSystem{rootPath: rootPath}
}

func (fs *LocalFileSystem) Read(filePath string) ([]byte, error) {
	fullPath := fs.rootPath + "/" + filePath
	bytes, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (fs *LocalFileSystem) Write(filePath string, base64Content string) error {
	fullPath := fs.rootPath + "/" + filePath
	content, err := base64.StdEncoding.DecodeString(base64Content)
	if err != nil {
		content = []byte(base64Content)
	}
	err = os.WriteFile(fullPath, content, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (fs *LocalFileSystem) Remove(filePath string) error {
	fullPath := fs.rootPath + "/" + filePath
	err := os.Remove(fullPath)
	if err != nil {
		return err
	}
	return nil
}
