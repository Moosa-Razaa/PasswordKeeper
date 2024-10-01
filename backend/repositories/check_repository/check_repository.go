package check_repository

import "backend/repositories"

func CheckRepository() bool {
	fileHandler := repositories.GetFileHandlerInstance()
	return fileHandler.CheckRepositoryExistence()
}
