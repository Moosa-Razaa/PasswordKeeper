package repositories

import (
	"os"
	"path/filepath"
	"sync"
)

type FileHandler struct {
	file  *os.File
	mutex sync.Mutex
}

var fileHandlerInstance *FileHandler
var once sync.Once

func GetFilePath() (string, error) {
	baseDirectory := os.Getenv("BASE_DIRECTORY")

	if baseDirectory == "" {
		baseDirectory = "/app/data"
	}

	filePath := filepath.Join(baseDirectory, "passwords.json")
	directory := filepath.Dir(filePath)

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err = os.MkdirAll(directory, 0755)
		if err != nil {
			return "", err
		}
	}

	return filePath, nil
}

func GetFileHandlerInstance() (*FileHandler, error) {
	var err error

	filePath, geFilePathError := GetFilePath()
	if geFilePathError != nil {
		return nil, geFilePathError
	}

	once.Do(func() {
		var file *os.File
		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return
		}
		fileHandlerInstance = &FileHandler{
			file: file,
		}
	})
	return fileHandlerInstance, nil
}

func (fileHandler *FileHandler) Write(data string) (int, error) {
	fileHandler.mutex.Lock()
	defer fileHandler.mutex.Unlock()

	return fileHandler.file.WriteString(data)
}

func (fileHandler *FileHandler) Read() (string, error) {
	fileHandler.mutex.Lock()
	defer fileHandler.mutex.Unlock()

	fileInfo, err := fileHandler.file.Stat()
	if err != nil {
		return "", err
	}

	data := make([]byte, fileInfo.Size())
	_, err = fileHandler.file.Read(data)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (fileHandler *FileHandler) Close() error {
	fileHandler.mutex.Lock()
	defer fileHandler.mutex.Unlock()

	return fileHandler.file.Close()
}
