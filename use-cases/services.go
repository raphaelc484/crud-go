package usecases

import (
	"crud-go/models"

	"github.com/google/uuid"
)

type ItemService interface {
	GetAllItems() ([]models.Item, error)
	CreateItem(item *models.Item) error
	GetItemByID(id uuid.UUID) (*models.Item, error)
	UpdateItem(id uuid.UUID, updatedItem *models.Item) (*models.Item, error)
	DeleteItem(id uuid.UUID) error
}
