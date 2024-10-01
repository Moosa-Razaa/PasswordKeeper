package controllers

import (
	"backend/services"
	"encoding/json"
	"net/http"
)

func GetAllPasswordsController(w http.ResponseWriter, r *http.Request) {
	response, httpStatusCode := services.GetAllPasswords()

	if httpStatusCode != http.StatusOK {
		w.WriteHeader(httpStatusCode)
		_, _ = w.Write([]byte("Cannot get all passwords"))
		return
	}

	w.WriteHeader(http.StatusOK)
	byteResponse, marshalError := json.Marshal(response)
	if marshalError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Cannot marshal response"))
		return
	}

	_, _ = w.Write(byteResponse)
}
