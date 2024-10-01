package services

import "backend/repositories/check_repository"

func CheckFileExistence() (string, int) {
	if check_repository.CheckRepository() {
		return "", 200
	} else {
		return "File does not exist", 404
	}
}
