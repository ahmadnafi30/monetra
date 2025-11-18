package main

import (
	"os"

	"github.com/ahmadnafi30/monetra/backend/Internal/handler/rest"
	"github.com/ahmadnafi30/monetra/backend/Internal/repository"
	"github.com/ahmadnafi30/monetra/backend/Internal/service"
	"github.com/ahmadnafi30/monetra/backend/model"
	"github.com/ahmadnafi30/monetra/backend/pkg/bcrypt"
	"github.com/ahmadnafi30/monetra/backend/pkg/jwt"
	"github.com/ahmadnafi30/monetra/backend/pkg/mailer"
	"github.com/ahmadnafi30/monetra/backend/pkg/middleware"
	"github.com/redis/go-redis/v9"
)

func main() {
	model.ConnectDatabase()
	db := model.DB

	bcryptService := bcrypt.Init()
	jwtAuth := jwt.Init()

	// Redis client menggunakan env vars
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "monetra_redis" // fallback
	}

	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: redisHost + ":" + redisPort,
	})

	// Mailer
	smtpMailer := mailer.NewSMTPMailer()

	// Repository
	repo := repository.NewRepository(db, redisClient)

	// Service
	svc := service.NewService(service.InitParam{
		Repository: repo,
		Bcrypt:     bcryptService,
		JwtAuth:    jwtAuth,
		Mailer:     smtpMailer,
	})

	// Middleware
	mw := middleware.Init(jwtAuth, svc)

	// REST
	r := rest.NewRest(svc, mw, jwtAuth)
	r.MountEndpoints()
	r.Run()
}
