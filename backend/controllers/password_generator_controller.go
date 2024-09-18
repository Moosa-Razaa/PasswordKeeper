package controllers

import (
	"backend/services"
	"net/http"
)

func PasswordGeneratorController(w http.ResponseWriter, r *http.Request) {
	response, httpStatusCode := services.GeneratePassword(r)

	if httpStatusCode != http.StatusOK {
		w.WriteHeader(httpStatusCode)
		_, _ = w.Write([]byte(response))
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(response))
}
