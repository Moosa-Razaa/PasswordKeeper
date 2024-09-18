package services

import (
	"backend/dtos/request"
	"backend/repositories"
	"backend/repositories/add_new_password_set"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"time"
)

func ExtractBodyAsPasswordRequest(body []byte) (request.PasswordRequest, error) {
	var req request.PasswordRequest
	err := json.Unmarshal(body, &req)
	if err != nil {
		return request.PasswordRequest{}, err
	}

	return req, nil
}

func ConvertRequestToPassword(req request.PasswordRequest) repositories.Password {
	return repositories.Password{
		PasswordSetId: uuid.New().String(),
		Email:         req.Email,
		Username:      req.Username,
		Password:      req.Password,
		CreatedAt:     time.Now().String(),
		UpdatedAt:     "",
		Domain:        req.Domain,
	}
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

	if requestBody.ValidatePasswordRequest() {
		return "Invalid request body", http.StatusBadRequest
	}

	password := ConvertRequestToPassword(requestBody)

	err := add_new_password_set.AddNewPasswordToRepository(password)
	if err != nil {
		return "Unable to add new password", http.StatusInternalServerError
	}

	return "Password added successfully", http.StatusOK
}
