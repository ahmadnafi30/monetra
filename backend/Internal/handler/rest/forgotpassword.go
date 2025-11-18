package rest

import (
	"net/http"

	"github.com/ahmadnafi30/monetra/backend/Internal/service"
	"github.com/gin-gonic/gin"
)

type OTPHandler struct {
	service service.OTPService
}

func NewOTPHandler(service service.OTPService) *OTPHandler {
	return &OTPHandler{service: service}
}

func (h *OTPHandler) RequestOTP(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.service.GenerateAndSendOTP(c.Request.Context(), req.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent"})
}

func (h *OTPHandler) VerifyOTP(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ok, err := h.service.VerifyOTP(c.Request.Context(), req.Email, req.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP verified"})
}

func (h *OTPHandler) ResetPassword(c *gin.Context) {
	var req struct {
		Email       string `json:"email"`
		Code        string `json:"code"`
		NewPassword string `json:"new_password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.NewPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "New password cannot be empty"})
		return
	}

	if len(req.NewPassword) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "New password must be at least 8 characters"})
		return
	}

	if err := h.service.ResetPassword(c.Request.Context(), req.Email, req.Code, req.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}

func (h *OTPHandler) ChangePassword(c *gin.Context) {
	var req struct {
		Email       string `json:"email"`
		NewPassword string `json:"new_password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.NewPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "New password cannot be empty"})
		return
	}

	if len(req.NewPassword) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "New password must be at least 8 characters"})
		return
	}

	if err := h.service.ChangePassword(c.Request.Context(), req.Email, req.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}
