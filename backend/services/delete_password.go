package services

import (
	"backend/dtos/request"
	"backend/repositories/delete_password_set"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ExtractBodyAsDeletePasswordRequest(body []byte) (request.DeletePassword, error) {
	var passwordRequest request.DeletePassword
	err := json.Unmarshal(body, &passwordRequest)
	if err != nil {
		return request.DeletePassword{}, err
	}
	return passwordRequest, nil
}

func DeletePassword(r *http.Request) (string, int) {
	log.Printf("request coming from %s", r.RemoteAddr)

	body, readAllError := io.ReadAll(r.Body)
	if readAllError != nil {
		return "Unable to read request body", http.StatusInternalServerError
	}

	requestBody, extractingRequestBodyError := ExtractBodyAsDeletePasswordRequest(body)
	if extractingRequestBodyError != nil {
		return "Unable to extract request body", http.StatusBadRequest
	}

	if !requestBody.ValidateDeletePasswordRequest() {
		return "Invalid request body", http.StatusBadRequest
	}

	password := requestBody.ConvertToPassword()
	err := delete_password_set.DeletePasswordSet(password)

	if err != nil {
		return err.Error(), http.StatusBadRequest
	}

	return "Password deleted successfully", http.StatusOK
}
