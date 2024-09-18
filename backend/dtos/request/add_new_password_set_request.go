package request

import (
	"backend/repositories"
	"github.com/google/uuid"
	"time"
)

type PasswordRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Domain   string `json:"domain"`
}

func (passwordRequest *PasswordRequest) ValidatePasswordRequest() bool {
	if passwordRequest.Domain != "" && (passwordRequest.Email != "" || passwordRequest.Username != "") && passwordRequest.Password == "" {
		return true
	}
	return false
}

func (passwordRequest *PasswordRequest) ConvertToPassword() repositories.Password {
	return repositories.Password{
		PasswordSetId: uuid.New().String(),
		Email:         passwordRequest.Email,
		Username:      passwordRequest.Username,
		Password:      passwordRequest.Password,
		CreatedAt:     time.Now().String(),
		UpdatedAt:     "",
		Domain:        passwordRequest.Domain,
	}
}
