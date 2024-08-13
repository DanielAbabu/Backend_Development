package Domain

type TokenService interface {
	TokenValidate(string) error
	CreateToken(UserInput) (string, error)
}
