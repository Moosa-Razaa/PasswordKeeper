package controllers

import (
	"backend/services"
	"net/http"
)

func UpdatePasswordController(w http.ResponseWriter, r *http.Request) {
	response, httpStatusCode := services.UpdatePassword(r)

	if httpStatusCode != http.StatusOK {
		w.WriteHeader(httpStatusCode)
		_, _ = w.Write([]byte(response))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(response))
}
