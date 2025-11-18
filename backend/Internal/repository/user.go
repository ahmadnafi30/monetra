package repository

import (
	"github.com/ahmadnafi30/monetra/backend/entity"
	"github.com/ahmadnafi30/monetra/backend/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	UpdatePassword(userID uuid.UUID, hash string) error
	GetUser(param model.UserParam) (entity.User, error)
	DeleteUser(userID uuid.UUID) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) UpdatePassword(userID uuid.UUID, hash string) error {
	return r.db.Model(&entity.User{}).
		Where("id = ?", userID).
		Update("password", hash).Error
}

func (r *UserRepository) GetUser(param model.UserParam) (entity.User, error) {
	var user entity.User
	query := r.db.Model(&entity.User{})

	if param.ID != uuid.Nil {
		query = query.Where("id = ?", param.ID)
	}

	if param.Email != "" {
		query = query.Where("email = ?", param.Email)
	}

	err := query.First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) DeleteUser(userID uuid.UUID) error {
	return r.db.Delete(&entity.User{}, "id = ?", userID).Error
}
