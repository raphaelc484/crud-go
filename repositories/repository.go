package repositories

import (
	"crud-go/models"

	"github.com/google/uuid"
)

type ItemRepository interface {
	GetAllItems() ([]models.Item, error)
	CreateItem(item *models.Item) error
	GetItemByID(id uuid.UUID) (*models.Item, error)
	UpdateItem(item *models.Item) error
	DeleteItem(id uuid.UUID) error
}
