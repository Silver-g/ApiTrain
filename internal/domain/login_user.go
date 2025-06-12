package domain

type LoginUserInternal struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	PasswordHash string `json:"hashpassword"`
	JwtToken     string `json:"authtoken"`
}
