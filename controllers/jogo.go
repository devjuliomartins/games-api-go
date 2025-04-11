package controllers

import "github.com/gin-gonic/gin"

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
