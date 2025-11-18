package rest

import (
	"context"
	"net/http"
	"os"

	"github.com/ahmadnafi30/monetra/backend/entity"
	"github.com/ahmadnafi30/monetra/backend/model"
	gconfig "github.com/ahmadnafi30/monetra/backend/pkg/oauth2"
	"github.com/ahmadnafi30/monetra/backend/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	goauth "golang.org/x/oauth2"
	"google.golang.org/api/idtoken"
)

func (r *Rest) Register(ctx *gin.Context) {
	param := model.UserRegister{}

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.ErrorCtx(ctx, http.StatusBadRequest, "Failed bind input", err)
		return
	}

	if err := r.service.UserService.Register(param); err != nil {
		response.ErrorCtx(ctx, http.StatusInternalServerError, "failed to register", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "success register new user", nil)
}

func (r *Rest) Login(ctx *gin.Context) {
	param := model.LoginAcc{}

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.ErrorCtx(ctx, http.StatusBadRequest, "invalid email or password", err)
		return
	}

	token, err := r.service.UserService.Login(param)
	if err != nil {
		response.ErrorCtx(ctx, http.StatusUnauthorized, "failed to login", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success login to system", token)
}

func (r *Rest) GoogleLogin(ctx *gin.Context) {
	conf := gconfig.GetGoogleOAuthConfig()
	url := conf.AuthCodeURL("random-state", goauth.AccessTypeOffline)
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (r *Rest) GoogleCallback(ctx *gin.Context) {
	conf := gconfig.GetGoogleOAuthConfig()
	code := ctx.Query("code")
	if code == "" {
		response.ErrorCtx(ctx, http.StatusBadRequest, "No code in query", nil)
		return
	}

	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		response.ErrorCtx(ctx, http.StatusInternalServerError, "Failed exchange token", err)
		return
	}

	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		response.ErrorCtx(ctx, http.StatusInternalServerError, "No id_token in token", nil)
		return
	}

	payload, err := idtoken.Validate(context.Background(), idToken, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		response.ErrorCtx(ctx, http.StatusUnauthorized, "Invalid ID token", err)
		return
	}

	email, ok := payload.Claims["email"].(string)
	if !ok {
		response.ErrorCtx(ctx, http.StatusInternalServerError, "Email not found in token", nil)
		return
	}

	name, _ := payload.Claims["name"].(string)

	user, err := r.service.UserService.GetUser(model.UserParam{Email: email})
	if err != nil {
		newUser := entity.User{
			ID:       uuid.New(),
			Email:    email,
			Name:     name,
			Password: "-",
		}
		_, err = r.service.UserService.CreateGoogleUser(newUser)
		if err != nil {
			response.ErrorCtx(ctx, http.StatusInternalServerError, "Failed to create user", err)
			return
		}
		user = newUser
	}

	jwtToken, err := r.service.UserService.GenerateToken(user.ID)
	if err != nil {
		response.ErrorCtx(ctx, http.StatusInternalServerError, "Failed create token", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success login with Google", model.UserLoginResponse{
		Token: jwtToken,
	})
}

func (r *Rest) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	userID, err := uuid.Parse(idParam)
	if err != nil {
		response.ErrorCtx(ctx, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	err = r.service.UserService.DeleteUser(userID)
	if err != nil {
		response.ErrorCtx(ctx, http.StatusInternalServerError, "Failed to delete user", err)
		return
	}

	response.Success(ctx, http.StatusOK, "User deleted successfully", nil)
}
