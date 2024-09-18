package add_new_password_set

import "backend/repositories"

func AddNewPasswordToRepository(password repositories.Password) error {
	fileHandler := repositories.GetFileHandlerInstance()

	doPasswordExist, passwordExistCheckError := fileHandler.CheckPasswordExists(password)

	if passwordExistCheckError != nil {
		return passwordExistCheckError
	}

	if !doPasswordExist {
		saveError := fileHandler.SaveNewPassword(password)
		if saveError != nil {
			return saveError
		}
	}

	return nil
}
