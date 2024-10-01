package services

import (
	"backend/repositories"
	"backend/repositories/get_all_password_set"
	"net/http"
)

func GetAllPasswords() ([]repositories.Password, int) {
	response, err := get_all_password_set.GetAllPasswords()
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	return response, http.StatusOK
}
