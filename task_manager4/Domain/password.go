package Domain

type PasswordService interface {
	HashPasword(string) (string, error)
	ComparePassword(string, string) (bool, error)
}
