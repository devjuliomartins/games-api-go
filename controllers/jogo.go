package controllers

import (
	"github.com/devjuliomartins/games-api-go/Paginacao/controller"
	"github.com/devjuliomartins/games-api-go/database"
	"github.com/devjuliomartins/games-api-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListarJogos(c *gin.Context) {
	var jogos []models.Jogo
	query := database.DB.Model(&models.Jogo{})

	meta, err := controller.PaginarConsulta(c, query, &jogos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"meta": meta,
		"data": jogos,
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

func VisualizarJogo(c *gin.Context) {
	var jogo models.Jogo
	id := c.Params.ByName("id")
	database.DB.First(&jogo, id)
	c.JSON(200, jogo)
}

func DeletarJogo(c *gin.Context) {
	var jogo models.Jogo
	id := c.Params.ByName("id")
	database.DB.Delete(&jogo, id)
	c.JSON(200, jogo)
}

func EditaJogo(c *gin.Context) {
	var jogo models.Jogo
	id := c.Params.ByName("id")
	database.DB.First(&jogo, id)
	if err := c.ShouldBind(&jogo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Model(&jogo).Updates(jogo)
	c.JSON(http.StatusOK, jogo)
}

func BuscarJogoPorGenero(c *gin.Context) {
	var jogos []models.Jogo
	genero := c.Param("genero")

	query := database.DB.Model(&models.Jogo{}).Where("genero LIKE ?", "%"+genero+"%")

	meta, err := controller.PaginarConsulta(c, query, &jogos)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"meta": meta,
		"data": jogos,
	})
}

func BuscarJogoPorPlataforma(c *gin.Context) {
	var jogos []models.Jogo
	plataforma := c.Param("plataforma")

	query := database.DB.Model(&models.Jogo{}).Where("plataforma LIKE ?", "%"+plataforma+"%")

	meta, err := controller.PaginarConsulta(c, query, &jogos)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"meta": meta,
		"data": jogos,
	})
}

func BuscarJogoPorDesenvolvedora(c *gin.Context) {
	var jogos []models.Jogo
	desenvolvedora := c.Param("desenvolvedora")

	query := database.DB.Model(&models.Jogo{}).Where("desenvolvedora LIKE ?", "%"+desenvolvedora+"%")

	meta, err := controller.PaginarConsulta(c, query, &jogos)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"meta": meta,
		"data": jogos,
	})
}

func BuscarJogoPorNota(c *gin.Context) {
	var jogos []models.Jogo
	nota := c.Param("nota")

	query := database.DB.Model(&models.Jogo{}).Where("nota >= ?", nota).Order("nota desc, titulo")

	meta, err := controller.PaginarConsulta(c, query, &jogos)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"meta": meta,
		"data": jogos,
	})
}
