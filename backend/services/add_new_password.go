package services

import (
	"backend/dtos/request"
	"backend/repositories/add_new_password_set"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ExtractBodyAsPasswordRequest(body []byte) (request.PasswordRequest, error) {
	var req request.PasswordRequest
	err := json.Unmarshal(body, &req)
	if err != nil {
		return request.PasswordRequest{}, err
	}

	return req, nil
}

func AddNewPassword(r *http.Request) (string, int) {
	log.Printf("request coming from %s", r.RemoteAddr)

	body, readAllError := io.ReadAll(r.Body)
	if readAllError != nil {
		return "Unable to read request body", http.StatusInternalServerError
	}

	requestBody, extractingRequestBodyError := ExtractBodyAsPasswordRequest(body)
	if extractingRequestBodyError != nil {
		return "Unable to extract request body", http.StatusBadRequest
	}

	if !requestBody.ValidatePasswordRequest() {
		return "Invalid request body", http.StatusBadRequest
	}

	password := requestBody.ConvertToPassword()

	err := add_new_password_set.AddNewPasswordToRepository(password)
	if err != nil {
		return "Unable to add new password", http.StatusInternalServerError
	}

	return "Password added successfully", http.StatusOK
}
