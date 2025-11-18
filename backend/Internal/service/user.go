package service

import (
	"fmt"

	"github.com/ahmadnafi30/monetra/backend/Internal/repository"
	"github.com/ahmadnafi30/monetra/backend/entity"
	"github.com/ahmadnafi30/monetra/backend/model"
	"github.com/ahmadnafi30/monetra/backend/pkg/bcrypt"
	"github.com/ahmadnafi30/monetra/backend/pkg/jwt"
	"github.com/google/uuid"
)

type UserService struct {
	user repository.IUserRepository
	// passwordResetRepo repository.IPasswordResetRepository
	bcrypt bcrypt.Interface
	jwt    jwt.Interface
}

type IUserService interface {
	Register(param model.UserRegister) error
	Login(param model.LoginAcc) (model.UserLoginResponse, error)
	GetUser(param model.UserParam) (entity.User, error)
	CreateGoogleUser(user entity.User) (entity.User, error)
	GenerateToken(userID uuid.UUID) (string, error)
	DeleteUser(userID uuid.UUID) error
}

func NewUserService(user repository.IUserRepository, bcrypt bcrypt.Interface, jwt jwt.Interface) IUserService {
	return &UserService{
		user:   user,
		bcrypt: bcrypt,
		jwt:    jwt,
	}
}

func (u *UserService) Register(param model.UserRegister) error {
	hashPassword, err := u.bcrypt.GenerateFromPassword(param.Password)
	if err != nil {
		return err
	}

	newUserID := uuid.New()

	user := entity.User{
		ID:       newUserID,
		Name:     param.Name,
		Email:    param.Email,
		Password: hashPassword,
		Provider: "manual",
	}

	_, err = u.user.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) GetUser(param model.UserParam) (entity.User, error) {
	return u.user.GetUser(param)
}

func (u *UserService) DeleteUser(userID uuid.UUID) error {
	return u.user.DeleteUser(userID)
}

func (u *UserService) Login(param model.LoginAcc) (model.UserLoginResponse, error) {
	var result model.UserLoginResponse

	user, err := u.user.GetUser(model.UserParam{
		Email: param.Email,
	})
	if err != nil {
		return result, err
	}

	if user.Provider == "google" {
		return result, fmt.Errorf("akun ini menggunakan login Google")
	}

	err = u.bcrypt.CompareHashAndPassword(user.Password, param.Password)
	if err != nil {
		return result, err
	}

	token, err := u.jwt.CreateToken(user.ID)
	if err != nil {
		return result, err
	}

	result.Token = token
	return result, nil
}

func (u *UserService) CreateGoogleUser(user entity.User) (entity.User, error) {
	user.Provider = "google"
	user.Password = ""

	createdUser, err := u.user.CreateUser(user)
	if err != nil {
		return entity.User{}, err
	}

	return createdUser, nil
}

func (u *UserService) GenerateToken(userID uuid.UUID) (string, error) {
	return u.jwt.CreateToken(userID)
}
