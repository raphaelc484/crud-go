package repositories

import (
	"crud-go/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db}
}

func (r *itemRepository) GetAllItems() ([]models.Item, error) {
	var items []models.Item
	err := r.db.Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *itemRepository) CreateItem(item *models.Item) error {
	item.ID = uuid.New()
	return r.db.Create(&item).Error
}

func (r *itemRepository) GetItemByID(id uuid.UUID) (*models.Item, error) {
	var item models.Item
	err := r.db.First(&item, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *itemRepository) UpdateItem(item *models.Item) error {
	return r.db.Save(item).Error
}

func (r *itemRepository) DeleteItem(item *models.Item) error {
	return r.db.Delete(&item).Error

}
