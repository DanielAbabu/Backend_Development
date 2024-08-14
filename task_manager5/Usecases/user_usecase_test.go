package Usecases_test

// import (
// 	"errors"
// 	"testing"

// 	"task_manager5/Domain"
// 	"task_manager5/Usecases"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// // MockUserRepository is a mock implementation of the UserRepository interface
// type MockUserRepository struct {
// 	mock.Mock
// }

// func (m *MockUserRepository) FindByEmail(email string) (Domain.DBUser, error) {
// 	args := m.Called(email)
// 	return args.Get(0).(Domain.DBUser), args.Error(1)
// }

// func (m *MockUserRepository) CreateUser(user Domain.UserInput) (Domain.DBUser, error) {
// 	args := m.Called(user)
// 	return args.Get(0).(Domain.DBUser), args.Error(1)
// }

// func (m *MockUserRepository) FindAllUsers() ([]Domain.DBUser, error) {
// 	args := m.Called()
// 	return args.Get(0).([]Domain.DBUser), args.Error(1)
// }

// func (m *MockUserRepository) FindById(id string) (Domain.DBUser, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(Domain.DBUser), args.Error(1)
// }

// func (m *MockUserRepository) UpdateUserById(id string, user Domain.UserInput, isAdmin bool) (Domain.DBUser, error) {
// 	args := m.Called(id, user, isAdmin)
// 	return args.Get(0).(Domain.DBUser), args.Error(1)
// }

// func (m *MockUserRepository) DeleteUserByID(id string) error {
// 	args := m.Called(id)
// 	return args.Error(0)
// }

// // MockPasswordService is a mock implementation of the PasswordService interface
// type MockPasswordService struct {
// 	mock.Mock
// }

// func (m *MockPasswordService) HashPasword(password string) (string, error) {
// 	args := m.Called(password)
// 	return args.String(0), args.Error(1)
// }

// func (m *MockPasswordService) ComparePassword(hashedPassword, password string) (bool, error) {
// 	args := m.Called(hashedPassword, password)
// 	return args.Bool(0), args.Error(1)
// }

// // MockTokenService is a mock implementation of the TokenService interface
// type MockTokenService struct {
// 	mock.Mock
// }

// func (m *MockTokenService) CreateToken(user Domain.DBUser) (string, error) {
// 	args := m.Called(user)
// 	return args.String(0), args.Error(1)
// }

// func TestLogin(t *testing.T) {
// 	mockRepo := new(MockUserRepository)
// 	mockPasswordService := new(MockPasswordService)
// 	mockTokenService := new(MockTokenService)
// 	userUC := Usecases.NewUserUC(mockRepo, mockPasswordService, mockTokenService)

// 	inputUser := Domain.UserInput{
// 		Email:    "test@example.com",
// 		Password: "password",
// 	}

// 	dbUser := Domain.DBUser{
// 		ID:       primitive.NewObjectID(),
// 		Email:    "test@example.com",
// 		Password: "hashed_password",
// 	}

// 	mockRepo.On("FindByEmail", inputUser.Email).Return(dbUser, nil)
// 	mockPasswordService.On("ComparePassword", dbUser.Password, inputUser.Password).Return(true, nil)
// 	mockTokenService.On("CreateToken", dbUser).Return("token", nil)

// 	resultUser, token, err := userUC.Login(inputUser)

// 	assert.NoError(t, err)
// 	assert.Equal(t, Domain.ChangeToOutput(dbUser), resultUser)
// 	assert.Equal(t, "token", token)
// 	mockRepo.AssertExpectations(t)
// 	mockPasswordService.AssertExpectations(t)
// 	mockTokenService.AssertExpectations(t)
// }

// func TestSignup(t *testing.T) {
// 	mockRepo := new(MockUserRepository)
// 	mockPasswordService := new(MockPasswordService)
// 	userUC := Usecases.NewUserUC(mockRepo, mockPasswordService, nil)

// 	inputUser := Domain.UserInput{
// 		Email:    "test@example.com",
// 		Password: "password",
// 	}

// 	hashedPassword := "hashed_password"
// 	dbUser := Domain.DBUser{
// 		ID:       primitive.NewObjectID(),
// 		Email:    "test@example.com",
// 		Password: hashedPassword,
// 	}

// 	mockPasswordService.On("HashPasword", inputUser.Password).Return(hashedPassword, nil)
// 	mockRepo.On("CreateUser", inputUser).Return(dbUser, nil)

// 	resultUser, err := userUC.Signup(inputUser)

// 	assert.NoError(t, err)
// 	assert.Equal(t, dbUser, resultUser)
// 	mockRepo.AssertExpectations(t)
// 	mockPasswordService.AssertExpectations(t)
// }

// func TestGetUsers(t *testing.T) {
// 	mockRepo := new(MockUserRepository)
// 	userUC := Usecases.NewUserUC(mockRepo, nil, nil)

// 	users := []Domain.DBUser{
// 		{Email: "user1@example.com"},
// 		{Email: "user2@example.com"},
// 	}

// 	mockRepo.On("FindAllUsers").Return(users, nil)

// 	result, err := userUC.GetUsers()

// 	assert.NoError(t, err)
// 	assert.Equal(t, users, result)
// 	mockRepo.AssertExpectations(t)
// }

// func TestGetUser(t *testing.T) {
// 	mockRepo := new(MockUserRepository)
// 	userUC := Usecases.NewUserUC(mockRepo, nil, nil)

// 	userID := "user-id"
// 	dbUser := Domain.DBUser{Email: "user@example.com"}

// 	mockRepo.On("FindById", userID).Return(dbUser, nil)

// 	result, err := userUC.GetUser(userID)

// 	assert.NoError(t, err)
// 	assert.Equal(t, Domain.ChangeToOutput(dbUser), result)
// 	mockRepo.AssertExpectations(t)
// }

// func TestMakeAdmin(t *testing.T) {
// 	mockRepo := new(MockUserRepository)
// 	userUC := Usecases.NewUserUC(mockRepo, nil, nil)

// 	userID := "user-id"
// 	dbUser := Domain.DBUser{Email: "user@example.com"}

// 	mockRepo.On("FindById", userID).Return(dbUser, nil)
// 	mockRepo.On("UpdateUserById", userID, mock.AnythingOfType("Domain.UserInput"), true).Return(dbUser, nil)

// 	result, err := userUC.MakeAdmin(userID)

// 	assert.NoError(t, err)
// 	assert.Equal(t, dbUser, result)
// 	mockRepo.AssertExpectations(t)
// }

// func TestUpdateUser(t *testing.T) {
// 	mockRepo := new(MockUserRepository)
// 	mockPasswordService := new(MockPasswordService)
// 	userUC := Usecases.NewUserUC(mockRepo, mockPasswordService, nil)

// 	userID := "user-id"
// 	inputUser := Domain.UserInput{
// 		Name:     "Updated Name",
// 		Email:    "updated@example.com",
// 		Password: "new_password",
// 	}

// 	dbUser := Domain.DBUser{
// 		ID:       primitive.NewObjectID(),
// 		Email:    "old@example.com",
// 		Password: "old_password",
// 	}

// 	hashedPassword := "hashed_new_password"
// 	updatedUser := Domain.DBUser{
// 		ID:       dbUser.ID,
// 		Name:     inputUser.Name,
// 		Email:    inputUser.Email,
// 		Password: hashedPassword,
// 	}

// 	mockRepo.On("FindById", userID).Return(dbUser, nil)
// 	mockPasswordService.On("HashPasword", inputUser.Password).Return(hashedPassword, nil)
// 	mockRepo.On("UpdateUserById", userID, mock.AnythingOfType("Domain.UserInput"), false).Return(updatedUser, nil)

// 	result, err := userUC.UpdateUser(userID, inputUser)

// 	assert.NoError(t, err)
// 	assert.Equal(t, updatedUser, result)
// 	mockRepo.AssertExpectations(t)
// 	mockPasswordService.AssertExpectations(t)
// }

// func TestDeleteUser(t *testing.T) {
// 	mockRepo := new(MockUserRepository)
// 	userUC := Usecases.NewUserUC(mockRepo, nil, nil)

// 	userID := "user-id"

// 	mockRepo.On("DeleteUserByID", userID).Return(nil)

// 	err := userUC.DeleteUser(userID)

// 	assert.NoError(t, err)
// 	mockRepo.AssertExpectations(t)
// }

// func TestLogin_InvalidPassword(t *testing.T) {
// 	mockRepo := new(MockUserRepository)
// 	mockPasswordService := new(MockPasswordService)
// 	mockTokenService := new(MockTokenService)
// 	userUC := Usecases.NewUserUC(mockRepo, mockPasswordService, mockTokenService)

// 	inputUser := Domain.UserInput{
// 		Email:    "test@example.com",
// 		Password: "wrong_password",
// 	}

// 	dbUser := Domain.DBUser{
// 		ID:       primitive.NewObjectID(),
// 		Email:    "test@example.com",
// 		Password: "hashed_password",
// 	}

// 	mockRepo.On("FindByEmail", inputUser.Email).Return(dbUser, nil)
// 	mockPasswordService.On("ComparePassword", dbUser.Password, inputUser.Password).Return(false, errors.New("invalid password"))

// 	_, _, err := userUC.Login(inputUser)

// 	assert.Error(t, err)
// 	assert.Equal(t, "invalid password", err.Error())
// 	mockRepo.AssertExpectations(t)
// 	mockPasswordService.AssertExpectations(t)
// }
