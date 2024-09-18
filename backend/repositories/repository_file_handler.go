package repositories

import (
	"encoding/json"
	"io"
	"os"
	"sync"
)

type FileHandler struct {
	filePath string
	mutex    sync.Mutex
}

func GetFileHandlerInstance() *FileHandler {
	var filePath = "passwords.json"
	return &FileHandler{filePath: filePath}
}

func (fileHandlerInstance *FileHandler) ReadAll() ([]Password, error) {
	fileHandlerInstance.mutex.Lock()
	defer fileHandlerInstance.mutex.Unlock()

	file, fileOpenError := os.Open(fileHandlerInstance.filePath)
	if fileOpenError != nil {
		return nil, fileOpenError
	}

	byteValue, readAllError := io.ReadAll(file)
	if readAllError != nil {
		return nil, readAllError
	}

	var passwords []Password
	unmarshalError := json.Unmarshal(byteValue, &passwords)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return passwords, nil
}

func (fileHandlerInstance *FileHandler) SaveNewPassword(newPassword Password) error {
	fileHandlerInstance.mutex.Lock()
	defer fileHandlerInstance.mutex.Unlock()

	file, fileOpenError := os.OpenFile(fileHandlerInstance.filePath, os.O_RDWR|os.O_CREATE, 0644)
	if fileOpenError != nil {
		return fileOpenError
	}

	allPasswords, readAllError := fileHandlerInstance.ReadAll()
	if readAllError != nil {
		return readAllError
	}

	allPasswords = append(allPasswords, newPassword)

	encodedPasswords, encodeError := json.Marshal(allPasswords)
	if encodeError != nil {
		return encodeError
	}

	_, writeError := file.Write(encodedPasswords)
	if writeError != nil {
		return writeError
	}

	return nil
}

func (fileHandlerInstance *FileHandler) CheckPasswordExists(password Password) (bool, error) {
	fileHandlerInstance.mutex.Lock()
	defer fileHandlerInstance.mutex.Unlock()

	allPasswords, readAllError := fileHandlerInstance.ReadAll()
	if readAllError != nil {
		return false, readAllError
	}

	for _, existingPassword := range allPasswords {
		if existingPassword.domain == password.domain && (existingPassword.Email == password.Email || existingPassword.Username == password.Username) {
			return true, nil
		}
	}

	return false, nil
}

func (fileHandlerInstance *FileHandler) GetPasswordIndex(password Password) (int, error) {
	fileHandlerInstance.mutex.Lock()
	defer fileHandlerInstance.mutex.Unlock()

	allPasswords, readAllError := fileHandlerInstance.ReadAll()
	if readAllError != nil {
		return -1, readAllError
	}

	for index, existingPassword := range allPasswords {
		if existingPassword.domain == password.domain && (existingPassword.Email == password.Email || existingPassword.Username == password.Username) {
			return index, nil
		}
	}

	return -1, nil
}

func (fileHandlerInstance *FileHandler) UpdatePassword(updatedPassword Password) error {
	fileHandlerInstance.mutex.Lock()
	defer fileHandlerInstance.mutex.Unlock()

	allPasswords, readAllError := fileHandlerInstance.ReadAll()
	if readAllError != nil {
		return readAllError
	}

	passwordIndex, getPasswordIndexError := fileHandlerInstance.GetPasswordIndex(updatedPassword)
	if getPasswordIndexError != nil {
		return getPasswordIndexError
	}

	allPasswords[passwordIndex] = updatedPassword

	encodedPasswords, encodeError := json.Marshal(allPasswords)
	if encodeError != nil {
		return encodeError
	}

	file, fileOpenError := os.OpenFile(fileHandlerInstance.filePath, os.O_RDWR|os.O_CREATE, 0644)
	if fileOpenError != nil {
		return fileOpenError
	}

	_, writeError := file.Write(encodedPasswords)
	if writeError != nil {
		return writeError
	}

	return nil
}

func (fileHandlerInstance *FileHandler) DeletePassword(password Password) error {
	fileHandlerInstance.mutex.Lock()
	defer fileHandlerInstance.mutex.Unlock()

	allPasswords, readAllError := fileHandlerInstance.ReadAll()
	if readAllError != nil {
		return readAllError
	}

	passwordIndex, getPasswordIndexError := fileHandlerInstance.GetPasswordIndex(password)
	if getPasswordIndexError != nil {
		return getPasswordIndexError
	}

	allPasswords = append(allPasswords[:passwordIndex], allPasswords[passwordIndex+1:]...)

	encodedPasswords, encodeError := json.Marshal(allPasswords)
	if encodeError != nil {
		return encodeError
	}

	file, fileOpenError := os.OpenFile(fileHandlerInstance.filePath, os.O_RDWR|os.O_CREATE, 0644)
	if fileOpenError != nil {
		return fileOpenError
	}

	_, writeError := file.Write(encodedPasswords)
	if writeError != nil {
		return writeError
	}

	return nil
}
