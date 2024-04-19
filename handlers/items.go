package handlers

import (
	"crud-go/models"
	usecases "crud-go/use-cases"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ItemHandler struct {
	service usecases.ItemService
}

func NewItemHandler(service usecases.ItemService) *ItemHandler {
	return &ItemHandler{service}
}

// @Summary Obter todos os itens
// @Description Retorna todos os itens na base de dados.
// @Produce json
// @Success 200 {array} models.Item
// @Router /items [get]
func (h *ItemHandler) GetItems(c *gin.Context) {
	items, err := h.service.GetAllItems()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, items)
}

// @Summary      Cria um novo item
// @Description  Cria um novo item com as informações fornecidas.
// @Param        name  body      string  true  "Nome do item" example({"name": "Novo item"})
// @Success      201   {object}  models.Item  "Item criado com sucesso"
// @Router       /items [post]
func (h *ItemHandler) CreateItem(c *gin.Context) {
	var item models.Item
	err := c.ShouldBindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.CreateItem(&item)
	if err != nil {
		return
	}
	c.JSON(http.StatusCreated, item)
}

// @Summary      Pega um item pelo ID
// @Description  Retorna o item com o ID especificado.
// @Param        id   path      string  true  "ID do item"
// @Success      200  {object}  models.Item  "Item retornado com sucesso"
// @Router       /items/{id} [get]
func (h *ItemHandler) GetItemByID(c *gin.Context) {
	id := c.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	item, err := h.service.GetItemByID(parsedID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

// @Summary      Atualiza um item pelo ID
// @Description  Atualiza um item com o ID especificado.
// @Param        id   path      string      true  "ID do item"
// @Param        name  body      string  true  "Nome do item" example({"name": "Trocar o nome"})
// @Success      200  {object}  models.Item  "Item atualizado com sucesso"
// @Router       /items/{id} [put]
func (h *ItemHandler) UpdateItem(c *gin.Context) {
	id := c.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var updatedItem models.Item
	err = c.ShouldBindJSON(&updatedItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.service.UpdateItem(parsedID, &updatedItem)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary      Deleta um item pelo ID
// @Description  Remove um item com o ID especificado.
// @Param        id   path      string  true  "ID do item"
// @Success      204  "Item removido com sucesso"
// @Router       /items/{id} [delete]
func (h *ItemHandler) DeleteItem(c *gin.Context) {
	id := c.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	err = h.service.DeleteItem(parsedID)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
		}
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
