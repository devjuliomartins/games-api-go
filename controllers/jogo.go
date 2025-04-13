package controllers

import (
	"github.com/devjuliomartins/games-api-go/database"
	"github.com/devjuliomartins/games-api-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExibirGames(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":             "1",
		"Titulo":         "CS:GO",
		"Genero":         "FPS",
		"AnoLancamento":  "2012",
		"Desenvolvedora": "Valve",
		"Nota":           "9.0",
	})
}

func CriarNovoJogo(c *gin.Context) {
	var jogo models.Jogo

	if err := c.ShouldBind(&jogo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&jogo)
	c.JSON(http.StatusOK, jogo)
}
