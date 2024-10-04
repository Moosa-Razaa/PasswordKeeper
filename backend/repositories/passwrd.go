package repositories

type Repository struct {
	Key             string
	DefaultPassword string
	Passwords       []Password
}

type Password struct {
	PasswordSetId string
	Email         string
	Username      string
	Password      string
	CreatedAt     string
	UpdatedAt     string
	Domain        string
}
