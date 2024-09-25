package request

import (
	"backend/repositories"
	"time"
)

type UpdatePasswordSet struct {
	Domain      string `json:"domain"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
}

func (updatePasswordSet *UpdatePasswordSet) ValidateUpdatePasswordRequest() bool {
	if updatePasswordSet.Domain == "" {
		return false
	}

	if updatePasswordSet.Username == "" && updatePasswordSet.Email == "" {
		return false
	}

	if updatePasswordSet.NewPassword == "" {
		return false
	}

	return true
}

func (updatePasswordSet *UpdatePasswordSet) ConvertToPassword() repositories.Password {
	return repositories.Password{
		Domain:    updatePasswordSet.Domain,
		Username:  updatePasswordSet.Username,
		Email:     updatePasswordSet.Email,
		Password:  updatePasswordSet.NewPassword,
		UpdatedAt: time.Now().String(),
	}
}
