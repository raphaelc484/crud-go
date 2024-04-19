package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Inicializa o Router utilizando as configs Default do gin
	router := gin.Default()

	//Initialize Routes
	initialize_routes(router)

	//para mudar a porta padrao é necessário utilizar o Run
	router.Run(":8080")
}
