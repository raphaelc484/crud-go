package router

import (
	"crud-go/config"
	handlers "crud-go/handlers"
	"crud-go/repositories"
	usecases "crud-go/use-cases"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initialize_routes(router *gin.Engine) {
	db := config.GetDB()

	repo := repositories.NewItemRepository(db)
	service := usecases.NewItemService(repo)
	handler := handlers.NewItemHandler(service)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/items", handler.GetItems)
	router.POST("/items", handler.CreateItem)
	router.GET("/items/:id", handler.GetItemByID)
	router.PUT("/items/:id", handler.UpdateItem)
	router.DELETE("/items/:id", handler.DeleteItem)

}
