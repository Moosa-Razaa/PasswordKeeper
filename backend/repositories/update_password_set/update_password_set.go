package update_password_set

import "backend/repositories"

func UpdateExistingPassword(updatePasswordSet repositories.Password) error {
	fileHandler := repositories.GetFileHandlerInstance()

	updateError := fileHandler.UpdatePassword(updatePasswordSet)
	if updateError != nil {
		return updateError
	}

	return nil
}
