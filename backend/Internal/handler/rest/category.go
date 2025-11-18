package rest

import (
	"net/http"

	"github.com/ahmadnafi30/monetra/backend/model"
	"github.com/ahmadnafi30/monetra/backend/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Rest) CreateCategory(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	var req model.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorCtx(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	err := r.service.CategoryService.CreateCategory(req.Name, req.Type, userID.(uuid.UUID).String(), req.Color, req.Icon)
	if err != nil {
		response.ErrorCtx(ctx, http.StatusInternalServerError, "Failed to create category", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "Category created successfully", nil)
}

func (r *Rest) GetCategories(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	categoryType := ctx.Query("type")

	param := model.CategoryParam{
		UserID: userID.(uuid.UUID),
		Type:   categoryType,
	}

	categories, err := r.service.CategoryService.GetAllCategory(param)
	if err != nil {
		response.ErrorCtx(ctx, http.StatusInternalServerError, "Failed to get categories", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Categories retrieved successfully", categories)
}

func (r *Rest) GetCategoryByID(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	categoryID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		response.ErrorCtx(ctx, http.StatusBadRequest, "Invalid category ID", err)
		return
	}

	category, err := r.service.CategoryService.GetByID(categoryID, userID.(uuid.UUID))
	if err != nil {
		response.ErrorCtx(ctx, http.StatusNotFound, "Category not found", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Category retrieved successfully", category)
}

func (r *Rest) UpdateCategory(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	categoryID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		response.ErrorCtx(ctx, http.StatusBadRequest, "Invalid category ID", err)
		return
	}

	var req model.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorCtx(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	err = r.service.CategoryService.UpdateCategory(categoryID, userID.(uuid.UUID), req)
	if err != nil {
		response.ErrorCtx(ctx, http.StatusInternalServerError, "Failed to update category", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Category updated successfully", nil)
}

func (r *Rest) DeleteCategory(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	categoryID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		response.ErrorCtx(ctx, http.StatusBadRequest, "Invalid category ID", err)
		return
	}

	err = r.service.CategoryService.DeleteCategory(categoryID, userID.(uuid.UUID))
	if err != nil {
		response.ErrorCtx(ctx, http.StatusInternalServerError, "Failed to delete category", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Category deleted successfully", nil)
}
