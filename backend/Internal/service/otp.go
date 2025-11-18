package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	// "github.com/ahmadnafim/LifeMate/model"
	// "github.com/ahmadnafim/LifeMate/pkg"
	"github.com/ahmadnafi30/monetra/backend/Internal/repository"
	"github.com/ahmadnafi30/monetra/backend/model"
	"github.com/ahmadnafi30/monetra/backend/pkg/bcrypt"
	"github.com/ahmadnafi30/monetra/backend/pkg/mailer"
)

type OTPService interface {
	GenerateAndSendOTP(ctx context.Context, email string) error
	VerifyOTP(ctx context.Context, email, code string) (bool, error)
	ResetPassword(ctx context.Context, email, code, newPassword string) error
	ChangePassword(ctx context.Context, email, newPassword string) error
}

type otpService struct {
	otpRepo  repository.IOTPRepository
	userRepo repository.IUserRepository
	bcrypt   bcrypt.Interface
	mailer   mailer.Mailer
}

func NewOTPService(
	otpRepo repository.IOTPRepository,
	userRepo repository.IUserRepository,
	bcrypt bcrypt.Interface,
	mailer mailer.Mailer,
) OTPService {
	return &otpService{
		otpRepo:  otpRepo,
		userRepo: userRepo,
		bcrypt:   bcrypt,
		mailer:   mailer,
	}
}

func (s *otpService) GenerateAndSendOTP(ctx context.Context, email string) error {
	otp, err := generateOTP()
	if err != nil {
		return fmt.Errorf("failed to generate OTP: %w", err)
	}

	// Simpan OTP ke Redis
	if err := s.otpRepo.Save(ctx, email, otp, 5*time.Minute); err != nil {
		return fmt.Errorf("failed to save OTP: %w", err)
	}

	// Email template HTML
	emailBody := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<body>
			<p>Hello,</p>
			<p>Your OTP code is:</p>
			<h2>%s</h2>
			<p>This code will expire in 5 minutes.</p>
		</body>
		</html>`, otp)

	// Kirim email
	if err := s.mailer.Send(email, "Your OTP Code", emailBody); err != nil {
		return fmt.Errorf("failed to send OTP email: %w", err)
	}

	return nil
}

func (s *otpService) VerifyOTP(ctx context.Context, email, code string) (bool, error) {
	stored, err := s.otpRepo.Get(ctx, email)
	if err != nil {
		return false, fmt.Errorf("failed to get OTP: %w", err)
	}
	if stored != code {
		return false, nil
	}

	if err := s.otpRepo.Delete(ctx, email); err != nil {
		return false, fmt.Errorf("failed to delete OTP: %w", err)
	}

	return true, nil
}

func (s *otpService) ResetPassword(ctx context.Context, email, code, newPassword string) error {
	stored, err := s.otpRepo.Get(ctx, email)
	if err != nil {
		return fmt.Errorf("OTP not found or expired")
	}
	if stored != code {
		return fmt.Errorf("invalid OTP")
	}

	if err := s.otpRepo.Delete(ctx, email); err != nil {
		return fmt.Errorf("failed to delete OTP: %w", err)
	}

	// hashed, err := s.bcrypt.GenerateFromPassword(newPassword)
	// if err != nil {
	// 	return fmt.Errorf("failed to hash password: %w", err)
	// }

	// user, err := s.userRepo.GetUser(model.UserParam{Email: email})
	// if err != nil {
	// 	return fmt.Errorf("user not found")
	// }

	// if err := s.userRepo.UpdatePassword(user.ID, hashed); err != nil {
	// 	return fmt.Errorf("failed to update password: %w", err)

	// }

	return nil
}

func (s *otpService) ChangePassword(ctx context.Context, email, newPassword string) error {
	hashed, err := s.bcrypt.GenerateFromPassword(newPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user, err := s.userRepo.GetUser(model.UserParam{Email: email})
	if err != nil {
		return fmt.Errorf("user not found")
	}

	if err := s.userRepo.UpdatePassword(user.ID, hashed); err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	return nil
}

func generateOTP() (string, error) {
	n := big.NewInt(1000000)
	num, err := rand.Int(rand.Reader, n)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", num.Int64()), nil
}
