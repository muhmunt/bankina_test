package repository

import (
	"go-technical-test-bankina/src/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindByID(ID int) (entity.User, error)
	DeleteUserByID(id int) (entity.User, error)
	FindAll(offset, limit int) ([]entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) FindAll(offset, limit int) ([]entity.User, error) {
	var users []entity.User

	err := r.db.Offset(offset).Limit(limit).Find(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *userRepository) Save(user entity.User) (entity.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByID(ID int) (entity.User, error) {
	var user entity.User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) DeleteUserByID(id int) (entity.User, error) {
	var deleteUser entity.User
	err := r.db.Where("id = ?", id).Delete(&deleteUser).Error

	if err != nil {
		return deleteUser, err
	}

	return deleteUser, nil

}
