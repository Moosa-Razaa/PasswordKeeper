package request

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
