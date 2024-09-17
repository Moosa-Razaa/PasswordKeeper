package repositories

type AddNewPasswordSetRequest struct {
	PasswordSetId string
	Email         string
	Username      string
	Password      string
	CreatedAt     string
	UpdatedAt     string
}
