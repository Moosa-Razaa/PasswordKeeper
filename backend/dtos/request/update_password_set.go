package request

import (
	"backend/repositories"
	"time"
)

type UpdatePasswordSet struct {
	Domain      string `json:"domain"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (updatePasswordSet *UpdatePasswordSet) ValidateUpdatePasswordRequest() bool {
	if updatePasswordSet.Domain != "" && (updatePasswordSet.Email != "" || updatePasswordSet.Username != "") && updatePasswordSet.OldPassword != "" && updatePasswordSet.NewPassword != "" {
		return true
	}

	return false
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
