package user

import (
	"errors"
	"go-technical-test-bankina/src/entity"
	"go-technical-test-bankina/src/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(request RegisterRequest) (entity.User, error)
	CheckEmailAvailable(email string) (bool, error)
	Login(request LoginRequest) (entity.User, error)
	GetUserByID(ID int) (entity.User, error)
	UpdateUser(ID UserIDRequest, request UserUpdateRequest) (entity.User, error)
	DeleteUserByID(ID UserIDRequest) (entity.User, error)
	FindUsers(offset, limit int) ([]entity.User, error)
}

type service struct {
	repository repository.UserRepository
}

func NewService(repository repository.UserRepository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(request RegisterRequest) (entity.User, error) {
	user := entity.User{}
	user.Name = request.Name
	user.Email = request.Email
	// user.CreatedAt = time.Now()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) CheckEmailAvailable(email string) (bool, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) Login(request LoginRequest) (entity.User, error) {
	email := request.Email
	password := request.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("Email not found.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetUserByID(ID int) (entity.User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found.")
	}

	return user, nil
}

func (s *service) FindUsers(offset, limit int) ([]entity.User, error) {
	users, err := s.repository.FindAll(offset, limit)

	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *service) UpdateUser(ID UserIDRequest, request UserUpdateRequest) (entity.User, error) {

	user, err := s.repository.FindByID(ID.ID)

	if err != nil {
		return user, err
	}

	if user.ID != ID.ID {
		return user, errors.New("Unauthorized user update.")
	}

	user.Name = request.Name
	user.Email = request.Email
	user.UpdatedAt = time.Now()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	updateUser, err := s.repository.Save(user)
	if err != nil {
		return updateUser, err
	}

	return updateUser, nil
}

func (s *service) DeleteUserByID(ID UserIDRequest) (entity.User, error) {
	user, err := s.repository.DeleteUserByID(ID.ID)

	if err != nil {
		return user, err
	}

	return user, nil
}
