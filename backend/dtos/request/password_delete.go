package request

import "backend/repositories"

type DeletePassword struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Domain   string `json:"domain"`
}

func (deletePassword *DeletePassword) ValidateDeletePasswordRequest() bool {
	if deletePassword.Domain != "" && (deletePassword.Email != "" || deletePassword.Username != "") {
		return true
	}
	return false
}

func (deletePassword *DeletePassword) ConvertToPassword() repositories.Password {
	return repositories.Password{
		Email:    deletePassword.Email,
		Username: deletePassword.Username,
		Domain:   deletePassword.Domain,
	}
}
