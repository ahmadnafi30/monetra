package service

import (
	"strings"

	"github.com/ahmadnafi30/monetra/backend/Internal/repository"
	"github.com/ahmadnafi30/monetra/backend/entity"
	"github.com/ahmadnafi30/monetra/backend/model"
	"github.com/ahmadnafi30/monetra/backend/pkg/response"
	"github.com/google/uuid"
)

type ICategoryService interface {
	CreateCategory(name string, categoryType string, userID string, color string, icon string) error
	GetAllCategory(param model.CategoryParam) ([]entity.Category, error)
	GetByID(id uuid.UUID, userID uuid.UUID) (*entity.Category, error)
	UpdateCategory(id uuid.UUID, userID uuid.UUID, req model.UpdateCategoryRequest) error
	DeleteCategory(id uuid.UUID, userID uuid.UUID) error
}

type CategoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) ICategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(name string, categoryType string, userID string, color string, icon string) error {
	name = strings.TrimSpace(name)
	categoryType = strings.TrimSpace(categoryType)
	color = strings.TrimSpace(color)
	icon = strings.TrimSpace(icon)

	if name == "" {
		return response.Error(400, "Category name cannot be empty", nil)
	}

	if categoryType == "" {
		return response.Error(400, "Category type cannot be empty", nil)
	}

	if categoryType != "income" && categoryType != "expense" {
		return response.Error(400, "Invalid category type", nil)
	}

	userUUID, err := uuid.Parse(strings.TrimSpace(userID))
	if err != nil {
		return response.Error(400, "Invalid user ID", err)
	}

	existingCat, err := s.repo.GetCategoryByName(userUUID, name)
	if err == nil && existingCat != nil {
		return response.Error(400, "Category name already exists", nil)
	}

	if icon == "" {
		icon = "default_icon"
	}

	if color == "" {
		color = "#a855f7"
	}

	category := &entity.Category{
		Name:   name,
		Type:   categoryType,
		UserID: userUUID,
		Icon:   icon,
		Color:  color,
	}

	if err := s.repo.CreateCategory(category); err != nil {
		return response.Error(500, "Failed to create category", err)
	}

	return nil
}

func (s *CategoryService) GetAllCategory(param model.CategoryParam) ([]entity.Category, error) {
	return s.repo.ListCategories(param.UserID, param.Type)
}

func (s *CategoryService) GetByID(id uuid.UUID, userID uuid.UUID) (*entity.Category, error) {
	category, err := s.repo.GetCategoryByID(id, userID)
	if err != nil {
		return nil, response.Error(404, "Category not found", err)
	}
	return category, nil
}

func (s *CategoryService) UpdateCategory(id uuid.UUID, userID uuid.UUID, req model.UpdateCategoryRequest) error {
	category, err := s.repo.GetCategoryByID(id, userID)
	if err != nil {
		return response.Error(404, "Category not found", err)
	}

	if req.Name != "" {
		category.Name = strings.TrimSpace(req.Name)
	}
	if req.Icon != "" {
		category.Icon = strings.TrimSpace(req.Icon)
	}
	if req.Color != "" {
		category.Color = strings.TrimSpace(req.Color)
	}

	if err := s.repo.UpdateCategory(category); err != nil {
		return response.Error(500, "Failed to update category", err)
	}

	return nil
}

func (s *CategoryService) DeleteCategory(id uuid.UUID, userID uuid.UUID) error {
	if err := s.repo.DeleteCategory(id, userID); err != nil {
		return response.Error(500, "Failed to delete category", err)
	}
	return nil
}
