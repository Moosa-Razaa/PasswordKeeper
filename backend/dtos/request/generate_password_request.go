package request

type PasswordType int

const (
	AlphabetsOnly PasswordType = iota
	AlphanumericsOnly
	NumericsOnly
)

type GeneratePasswordRequest struct {
	Length                   int          `json:"length"`
	Type                     PasswordType `json:"type"`
	IncludeSpecialCharacters bool         `json:"include_special_characters"`
	SpecialCharacters        string       `json:"special_characters"`
}
