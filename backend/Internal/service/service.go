package service

import (
	"github.com/ahmadnafi30/monetra/backend/Internal/repository"
	"github.com/ahmadnafi30/monetra/backend/pkg/bcrypt"
	"github.com/ahmadnafi30/monetra/backend/pkg/jwt"
	"github.com/ahmadnafi30/monetra/backend/pkg/mailer"
)

type Service struct {
	UserService     IUserService
	OTPService      OTPService
	CategoryService ICategoryService
}

type InitParam struct {
	Repository *repository.Repository
	Bcrypt     bcrypt.Interface
	JwtAuth    jwt.Interface
	Mailer     mailer.Mailer
}

func NewService(param InitParam) *Service {
	userService := NewUserService(param.Repository.User, param.Bcrypt, param.JwtAuth)
	otpService := NewOTPService(param.Repository.OTP, param.Repository.User, param.Bcrypt, param.Mailer)
	categoryService := NewCategoryService(param.Repository.Category)

	return &Service{
		UserService:     userService,
		OTPService:      otpService,
		CategoryService: categoryService,
	}
}
