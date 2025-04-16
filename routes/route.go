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
	r.GET("/games/genero/:genero", controllers.BuscarJogoPorGenero)
	r.GET("/games/plataforma/:plataforma", controllers.BuscarJogoPorPlataforma)
	r.GET("/games/desenvolvedora/:desenvolvedora", controllers.BuscarJogoPorDesenvolvedora)

	r.Run(":8080")
}
