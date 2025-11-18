package rest

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ahmadnafi30/monetra/backend/Internal/service"
	"github.com/ahmadnafi30/monetra/backend/pkg/jwt"
	"github.com/ahmadnafi30/monetra/backend/pkg/middleware"
	"github.com/ahmadnafi30/monetra/backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	middleware middleware.Interface
	jwt        jwt.Interface
}

func NewRest(service *service.Service, middleware middleware.Interface, jwt jwt.Interface) *Rest {
	return &Rest{
		router:     gin.Default(),
		service:    service,
		middleware: middleware,
		jwt:        jwt,
	}
}

func (r *Rest) MountEndpoints() {
	r.router.Use(r.middleware.Timeout())
	r.router.Use(middleware.Cors())

	api := r.router.Group("/api/v1")

	authGroup := api.Group("/auth")
	{
		authGroup.POST("/register", r.Register)
		authGroup.POST("/login", r.Login)
		authGroup.GET("/google/login", r.GoogleLogin)
		authGroup.GET("/google/callback", r.GoogleCallback)
	}

	usersGroup := api.Group("/users")
	usersGroup.Use(r.middleware.AuthenticateUser)
	{
		// usersGroup.GET("/profile", r.GetUserProfile)
		// usersGroup.PUT("/profile", r.UpdateUserProfile)
		usersGroup.DELETE("/:id", r.DeleteUser)
	}

	otpHandler := NewOTPHandler(r.service.OTPService)
	otpGroup := api.Group("/otp")
	{
		otpGroup.POST("/request", otpHandler.RequestOTP)
		otpGroup.POST("/verify", otpHandler.VerifyOTP)
		otpGroup.POST("/reset-password", otpHandler.ResetPassword)
		otpGroup.POST("/change-password", otpHandler.ChangePassword)
	}

	categoriesGroup := api.Group("/categories")
	categoriesGroup.Use(r.middleware.AuthenticateUser)
	{
		categoriesGroup.POST("", r.CreateCategory)
		categoriesGroup.GET("", r.GetCategories)
		categoriesGroup.GET("/:id", r.GetCategoryByID)
		categoriesGroup.PUT("/:id", r.UpdateCategory)
		categoriesGroup.DELETE("/:id", r.DeleteCategory)
	}
}

func (r *Rest) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	r.router.Run(fmt.Sprintf(":%s", port))
}

func (r *Rest) TestTimeout(ctx *gin.Context) {
	time.Sleep(3 * time.Second)
	response.Success(ctx, http.StatusOK, "success", nil)
}
