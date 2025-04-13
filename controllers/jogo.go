package controllers

import (
	"github.com/devjuliomartins/games-api-go/database"
	"github.com/devjuliomartins/games-api-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListarJogos(c *gin.Context) {
	var jogo []models.Jogo
	database.DB.Find(&jogo)
	c.JSON(200, jogo)
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
