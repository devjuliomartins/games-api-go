package routes

import (
	"github.com/devjuliomartins/games-api-go/controllers"
	"github.com/gin-gonic/gin"
)

func RouteGames() {
	r := gin.Default()

	r.GET("/games", controllers.ExibirGames)
	r.POST("/games", controllers.CriarNovoJogo)

	r.Run(":8080")
}
