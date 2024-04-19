package usecases

import (
	"crud-go/models"
	"crud-go/repositories"

	"github.com/google/uuid"
)

type itemService struct {
	repo repositories.ItemRepository
}

func NewItemService(repo repositories.ItemRepository) ItemService {
	return &itemService{repo}
}

func (s *itemService) GetAllItems() ([]models.Item, error) {
	return s.repo.GetAllItems()
}

func (s *itemService) CreateItem(item *models.Item) error {
	return s.repo.CreateItem(item)
}

func (s *itemService) GetItemByID(id uuid.UUID) (*models.Item, error) {
	return s.repo.GetItemByID(id)
}

func (s *itemService) UpdateItem(id uuid.UUID, updatedItem *models.Item) (*models.Item, error) {
	// Obter o item existente pelo ID
	existingItem, err := s.repo.GetItemByID(id)
	if err != nil {
		return nil, err
	}

	// Atualizar o item existente com os dados do item atualizado
	existingItem.Name = updatedItem.Name

	// Salvar o item atualizado no repositório
	err = s.repo.UpdateItem(existingItem)
	if err != nil {
		return nil, err
	}

	return existingItem, nil
}

func (s *itemService) DeleteItem(id uuid.UUID) error {
	return s.repo.DeleteItem(id)
}
