package delete_password_set

import "backend/repositories"

func DeletePasswordSet(password repositories.Password) error {
	fileHandler := repositories.GetFileHandlerInstance()

	deleteError := fileHandler.DeletePassword(password)

	if deleteError != nil {
		return deleteError
	}

	return nil
}
