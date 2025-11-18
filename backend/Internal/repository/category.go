package repository

import (
	"github.com/ahmadnafi30/monetra/backend/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(category *entity.Category) error
	GetCategoryByID(id uuid.UUID, userID uuid.UUID) (*entity.Category, error)
	UpdateCategory(category *entity.Category) error
	DeleteCategory(id uuid.UUID, userID uuid.UUID) error
	ListCategories(userID uuid.UUID, categoryType string) ([]entity.Category, error)
	SortCatgoriesbyAlpabet(userID uuid.UUID, categoryType string) ([]entity.Category, error)
	GetCategoryByType(userID uuid.UUID, categoryType string) ([]entity.Category, error)
	GetCategoryByName(userID uuid.UUID, name string) (*entity.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}
func (r *categoryRepository) CreateCategory(category *entity.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) GetCategoryByID(id uuid.UUID, userID uuid.UUID) (*entity.Category, error) {
	var category entity.Category
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) GetCategoryByType(userID uuid.UUID, categoryType string) ([]entity.Category, error) {
	var categories []entity.Category
	err := r.db.Where("user_id = ? AND type = ?", userID, categoryType).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) GetCategoryByName(userID uuid.UUID, name string) (*entity.Category, error) {
	var category entity.Category
	err := r.db.Where("user_id = ? AND name = ?", userID, name).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) UpdateCategory(category *entity.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) DeleteCategory(id uuid.UUID, userID uuid.UUID) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&entity.Category{}).Error
}

func (r *categoryRepository) ListCategories(userID uuid.UUID, categoryType string) ([]entity.Category, error) {
	var categories []entity.Category
	query := r.db.Where("user_id = ?", userID)
	if categoryType != "" {
		query = query.Where("type = ?", categoryType)
	}
	err := query.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) SortCatgoriesbyAlpabet(userID uuid.UUID, categoryType string) ([]entity.Category, error) {
	var categories []entity.Category
	query := r.db.Where("user_id = ?", userID)
	if categoryType != "" {
		query = query.Where("type = ?", categoryType)
	}
	err := query.Order("name asc").Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) IsCategoryExist(name string, userID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&entity.Category{}).
		Where("name = ? AND user_id = ?", name, userID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *categoryRepository) SeedDefaultCategories(userID uuid.UUID) error {
	defaultCategories := []entity.Category{
		{ID: uuid.New(), UserID: userID, Name: "Gaji", Type: "income", Icon: "💰", Color: "#10b981", IsDefault: true},
		{ID: uuid.New(), UserID: userID, Name: "Bonus", Type: "income", Icon: "🎁", Color: "#3b82f6", IsDefault: true},
		{ID: uuid.New(), UserID: userID, Name: "Investasi", Type: "income", Icon: "📈", Color: "#8b5cf6", IsDefault: true},
		{ID: uuid.New(), UserID: userID, Name: "Makanan", Type: "expense", Icon: "🍔", Color: "#ef4444", IsDefault: true},
		{ID: uuid.New(), UserID: userID, Name: "Transport", Type: "expense", Icon: "🚗", Color: "#f59e0b", IsDefault: true},
		{ID: uuid.New(), UserID: userID, Name: "Belanja", Type: "expense", Icon: "🛒", Color: "#ec4899", IsDefault: true},
		{ID: uuid.New(), UserID: userID, Name: "Tagihan", Type: "expense", Icon: "📱", Color: "#6366f1", IsDefault: true},
		{ID: uuid.New(), UserID: userID, Name: "Hiburan", Type: "expense", Icon: "🎮", Color: "#14b8a6", IsDefault: true},
	}

	for _, cat := range defaultCategories {
		var count int64
		if err := r.db.Model(&entity.Category{}).
			Where("user_id = ? AND name = ?", userID, cat.Name).
			Count(&count).Error; err != nil {
			return err
		}

		if count == 0 {
			if err := r.db.Create(&cat).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
