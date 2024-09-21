package repositories

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

func InitializeFileIfNotExist() error {
	log.Printf("Checking if file exists")

	var filePath = "passwords.json"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		_, fileCreationError := os.Create(filePath)
		return fileCreationError
	}

	return nil
}

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

	if len(byteValue) == 0 {
		return []Password{}, nil
	}

	var passwords []Password
	unmarshalError := json.Unmarshal(byteValue, &passwords)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return passwords, nil
}

func (fileHandlerInstance *FileHandler) SaveNewPassword(newPassword Password) error {
	allPasswords, readAllError := fileHandlerInstance.ReadAll()
	if readAllError != nil {
		return readAllError
	}

	fileHandlerInstance.mutex.Lock()
	defer fileHandlerInstance.mutex.Unlock()

	file, fileOpenError := os.OpenFile(fileHandlerInstance.filePath, os.O_RDWR|os.O_CREATE, 0644)
	if fileOpenError != nil {
		return fileOpenError
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
	allPasswords, readAllError := fileHandlerInstance.ReadAll()
	if readAllError != nil {
		return false, readAllError
	}

	for _, existingPassword := range allPasswords {
		if existingPassword.Domain == password.Domain && (existingPassword.Email == password.Email || existingPassword.Username == password.Username) {
			return true, nil
		}
	}

	return false, nil
}

func (fileHandlerInstance *FileHandler) GetPasswordIndex(password Password) (int, error) {
	allPasswords, readAllError := fileHandlerInstance.ReadAll()
	if readAllError != nil {
		return -1, readAllError
	}

	if password.PasswordSetId == "" {
		for index, existingPassword := range allPasswords {
			if existingPassword.Domain == password.Domain && (existingPassword.Email == password.Email || existingPassword.Username == password.Username) {
				return index, nil
			}
		}
	} else {
		for index, existingPassword := range allPasswords {
			if existingPassword.PasswordSetId == password.PasswordSetId {
				return index, nil
			}
		}
	}

	return -1, errors.New("item not found")
}

func (fileHandlerInstance *FileHandler) UpdatePassword(updatedPassword Password) error {
	allPasswords, readAllError := fileHandlerInstance.ReadAll()
	if readAllError != nil {
		return readAllError
	}

	passwordIndex, getPasswordIndexError := fileHandlerInstance.GetPasswordIndex(updatedPassword)
	if getPasswordIndexError != nil {
		return getPasswordIndexError
	}

	previousPassword := allPasswords[passwordIndex]
	updatedPassword.PasswordSetId = previousPassword.PasswordSetId
	updatedPassword.CreatedAt = previousPassword.CreatedAt
	updatedPassword.UpdatedAt = time.Now().String()

	allPasswords[passwordIndex] = updatedPassword

	encodedPasswords, encodeError := json.Marshal(allPasswords)
	if encodeError != nil {
		return encodeError
	}

	fileHandlerInstance.mutex.Lock()
	defer fileHandlerInstance.mutex.Unlock()

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

	fileHandlerInstance.mutex.Lock()
	defer fileHandlerInstance.mutex.Unlock()

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
