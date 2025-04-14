package routes

import (
	"github.com/devjuliomartins/games-api-go/controllers"
	"github.com/gin-gonic/gin"
)

func RouteGames() {
	r := gin.Default()

	r.GET("/games", controllers.ListarJogos)
	r.POST("/games", controllers.CriarNovoJogo)
	r.GET("/games/:id", controllers.VisualizarJogo)
	r.DELETE("/games/:id", controllers.DeletarJogo)
	r.PATCH("/games/:id", controllers.EditaJogo)

	r.Run(":8080")
}
