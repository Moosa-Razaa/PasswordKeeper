package request

type AddNewPasswordSetRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
