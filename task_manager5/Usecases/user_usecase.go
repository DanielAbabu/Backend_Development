package Usecases

import "task_manager5/Domain"

type UserUC struct {
	repo      Domain.UserRepository
	PasswordS Domain.PasswordService
	TokenS    Domain.TokenService
}

func NewUserUC(repository Domain.UserRepository, PS Domain.PasswordService, TS Domain.TokenService) *UserUC {
	return &UserUC{
		repo:      repository,
		PasswordS: PS,
		TokenS:    TS,
	}
}

func (UUC *UserUC) Login(user Domain.UserInput) (Domain.DBUser, string, error) {
	usr, err := UUC.repo.FindByEmail(user.Email)
	if err != nil {
		return Domain.DBUser{}, "", err
	}

	_, er := UUC.PasswordS.ComparePassword(usr.Password, user.Password)
	if er != nil {
		return Domain.DBUser{}, "", er
	}

	token, terr := UUC.TokenS.CreateToken(usr)
	if terr != nil {
		return Domain.DBUser{}, "", terr
	}

	return Domain.ChangeToOutput(usr), token, nil

}
func (UUC *UserUC) Signup(user Domain.UserInput) (Domain.DBUser, error) {
	hashed_password, err := UUC.PasswordS.HashPasword(user.Password)
	if err != nil {
		return Domain.DBUser{}, err
	}

	user.Password = hashed_password
	usr, er := UUC.repo.CreateUser(user)

	if er != nil {
		return Domain.DBUser{}, er
	}

	return usr, nil

}
func (UUC *UserUC) GetUsers() ([]Domain.DBUser, error) {
	return UUC.repo.FindAllUsers()
}
func (UUC *UserUC) GetUser(id string) (Domain.DBUser, error) {
	user, err := UUC.repo.FindById(id)
	return Domain.ChangeToOutput(user), err
}
func (UUC *UserUC) MakeAdmin(id string) (Domain.DBUser, error) {
	user, err := UUC.repo.FindById(id)
	if err != nil {
		return Domain.DBUser{}, err
	}
	return UUC.repo.UpdateUserById(id, user, true)
}
func (UUC *UserUC) UpdateUser(id string, user Domain.UserInput) (Domain.DBUser, error) {
	usr, err := UUC.repo.FindById(id)
	if err != nil {
		return Domain.DBUser{}, err
	}
	if user.Password != "" {
		hashed_password, err := UUC.PasswordS.HashPasword(user.Password)
		if err != nil {
			return Domain.DBUser{}, err
		}
		user.Password = hashed_password
	} else {
		user.Password = usr.Password
	}

	if user.Name == "" {
		user.Name = usr.Name
	}
	if user.Email == "" {
		user.Email = usr.Email
	}

	return UUC.repo.UpdateUserById(id, user, false)
}
func (UUC *UserUC) DeleteUser(id string) error {
	return UUC.repo.DeleteUserByID(id)
}
