package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type Item struct {
	BaseModel
	Name string `json:"name"`
}

type ItemRequest struct {
	Name string `json:"name"`
}

func (dto *ItemRequest) Validate() error {
	if dto.Name == "" {
		return fmt.Errorf("o campo Name é obrigatório")
	}
	return nil
}

func ToCreateItem(dto *ItemRequest) *Item {
	return &Item{
		BaseModel: BaseModel{
			ID: uuid.New(),
		},
		Name: dto.Name,
	}
}

func ToUpdateItem(id uuid.UUID, dto *ItemRequest) *Item {
	return &Item{
		BaseModel: BaseModel{
			ID: id,
		},
		Name: dto.Name,
	}
}
