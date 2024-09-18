package services

import (
	"backend/repositories/delete_password_set"
	"io"
	"log"
	"net/http"
)

func DeletePassword(r *http.Request) (string, int) {
	log.Printf("request coming from %s", r.RemoteAddr)

	body, readAllError := io.ReadAll(r.Body)
	if readAllError != nil {
		return "Unable to read request body", http.StatusInternalServerError
	}

	requestBody, extractingRequestBodyError := ExtractBodyAsPasswordRequest(body)
	if extractingRequestBodyError != nil {
		return "Unable to extract request body", http.StatusBadRequest
	}

	if requestBody.ValidatePasswordRequest() {
		return "Invalid request body", http.StatusBadRequest
	}

	password := requestBody.ConvertToPassword()
	err := delete_password_set.DeletePasswordSet(password)

	if err != nil {
		return "Unable to delete password", http.StatusInternalServerError
	}

	return "Password deleted successfully", http.StatusOK
}
