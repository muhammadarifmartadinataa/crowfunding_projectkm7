package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// Mock repository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Save(user User) (User, error) {
	args := m.Called(user)
	return args.Get(0).(User), args.Error(1)
}

func (m *MockRepository) FindByEmail(email string) (User, error) {
	args := m.Called(email)
	return args.Get(0).(User), args.Error(1)
}

func (m *MockRepository) FindById(ID int) (User, error) {
	args := m.Called(ID)
	return args.Get(0).(User), args.Error(1)
}

func (m *MockRepository) Update(user User) (User, error) {
	args := m.Called(user)
	return args.Get(0).(User), args.Error(1)
}

// Test RegisterUser
func TestRegisterUser(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	input := RegisterUserInput{
		Name:       "John Doe",
		Email:      "johndoe@example.com",
		Occupation: "Developer",
		Password:   "password",
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	mockUser := User{
		ID:           1,
		Name:         input.Name,
		Email:        input.Email,
		Occupation:   input.Occupation,
		PasswordHash: string(hashedPassword),
		Role:         "user",
	}

	mockRepo.On("Save", mock.Anything).Return(mockUser, nil)

	result, err := service.RegisterUser(input)

	assert.NoError(t, err)
	assert.Equal(t, mockUser.ID, result.ID)
	assert.Equal(t, mockUser.Name, result.Name)
	assert.Equal(t, mockUser.Email, result.Email)
	mockRepo.AssertExpectations(t)
}

// Test Login
func TestLogin(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	input := LoginInput{
		Email:    "johndoe@example.com",
		Password: "password",
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	mockUser := User{
		ID:           1,
		Email:        input.Email,
		PasswordHash: string(hashedPassword),
	}

	mockRepo.On("FindByEmail", input.Email).Return(mockUser, nil)

	result, err := service.Login(input)

	assert.NoError(t, err)
	assert.Equal(t, mockUser.ID, result.ID)
	assert.Equal(t, mockUser.Email, result.Email)
	mockRepo.AssertExpectations(t)
}

// Test IsEmailAvailable
func TestIsEmailAvailable(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	input := CheckEmailInput{
		Email: "johndoe@example.com",
	}

	mockRepo.On("FindByEmail", input.Email).Return(User{}, nil)

	result, err := service.IsEmailAvailable(input)

	assert.NoError(t, err)
	assert.True(t, result)
	mockRepo.AssertExpectations(t)
}
