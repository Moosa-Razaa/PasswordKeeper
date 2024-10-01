package get_all_password_set

import "backend/repositories"

func GetAllPasswords() ([]repositories.Password, error) {
	fileHandler := repositories.GetFileHandlerInstance()
	passwords, readError := fileHandler.ReadAll()

	if readError != nil {
		return nil, readError
	}

	return passwords, nil
}
