package middleware

import (
	"github.com/ahmadnafi30/monetra/backend/pkg/jwt"

	"github.com/ahmadnafi30/monetra/backend/Internal/service"
	"github.com/gin-gonic/gin"
)

type Interface interface {
	Timeout() gin.HandlerFunc
	AuthenticateUser(ctx *gin.Context)
}

type middleware struct {
	jwtAuth jwt.Interface
	service *service.Service
}

func Init(jwtAuth jwt.Interface, service *service.Service) Interface {
	return &middleware{
		jwtAuth: jwtAuth,
		service: service,
	}
}
