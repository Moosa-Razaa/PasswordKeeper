package services

import (
	"backend/dtos/request"
	"backend/repositories/update_password_set"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ExtractBodyAsUpdatePasswordRequest(body []byte) (request.UpdatePasswordSet, error) {
	var req request.UpdatePasswordSet
	err := json.Unmarshal(body, &req)
	if err != nil {
		return request.UpdatePasswordSet{}, err
	}

	return req, nil
}

func UpdatePassword(r *http.Request) (string, int) {
	log.Printf("request coming from %s", r.RemoteAddr)

	body, readAllError := io.ReadAll(r.Body)
	if readAllError != nil {
		return "error reading request body", http.StatusInternalServerError
	}

	requestBody, extractingRequestBodyError := ExtractBodyAsUpdatePasswordRequest(body)
	if extractingRequestBodyError != nil {
		return "error extracting request body", http.StatusBadRequest
	}

	if requestBody.ValidateUpdatePasswordRequest() {
		return "invalid request body", http.StatusBadRequest
	}

	password := requestBody.ConvertToPassword()
	err := update_password_set.UpdateExistingPassword(password)

	if err != nil {
		return "error updating password", http.StatusInternalServerError
	}

	return "password updated", http.StatusOK
}
