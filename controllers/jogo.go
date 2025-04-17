package controllers

import (
	"fmt"
	"github.com/devjuliomartins/games-api-go/database"
	"github.com/devjuliomartins/games-api-go/models"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

type PaginationMeta struct {
	TotalJogos   int `json:"total_jogos"`
	Limite       int `json:"limite"`
	PaginaAtual  int `json:"pagina_atual"`
	TotalPaginas int `json:"total_paginas"`
}

func GetPaginationParams(c *gin.Context) (int, int, int, error) {
	paginaStr := c.DefaultQuery("pagina", "1")
	limiteStr := c.DefaultQuery("limite", "10")

	pagina, err := strconv.Atoi(paginaStr)
	if err != nil || pagina < 1 {
		return 0, 0, 0, fmt.Errorf("invalid 'pagina' parameter")
	}

	limite, err := strconv.Atoi(limiteStr)
	if err != nil || limite < 1 {
		return 0, 0, 0, fmt.Errorf("invalid 'limite' parameter")
	}

	offset := (pagina - 1) * limite
	return pagina, limite, offset, nil
}

func ListarJogos(c *gin.Context) {
	pagina, limite, offset, err := GetPaginationParams(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var total_jogos int64
	if err := database.DB.Model(&models.Jogo{}).Count(&total_jogos).Error; err != nil {
		c.JSON(500, gin.H{"error": "error counting games"})
		return
	}

	var jogos []models.Jogo
	if err := database.DB.Offset(offset).Limit(limite).Find(&jogos).Error; err != nil {
		c.JSON(500, gin.H{"error": "error fetching games"})
		return
	}

	total_paginas := int(math.Ceil(float64(total_jogos) / float64(limite)))

	c.JSON(200, gin.H{
		"meta": PaginationMeta{
			TotalJogos:   int(total_jogos),
			Limite:       limite,
			PaginaAtual:  pagina,
			TotalPaginas: total_paginas,
		},
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
	database.DB.Where("genero LIKE ?", "%"+genero+"%").Find(&jogos)
	c.JSON(200, jogos)
}

func BuscarJogoPorPlataforma(c *gin.Context) {
	var jogos []models.Jogo
	plataforma := c.Param("plataforma")
	database.DB.Where("plataforma LIKE ?", "%"+plataforma+"%").Find(&jogos)
	c.JSON(200, jogos)
}

func BuscarJogoPorDesenvolvedora(c *gin.Context) {
	var jogos []models.Jogo
	desenvolvedora := c.Param("desenvolvedora")
	database.DB.Where("desenvolvedora LIKE ?", "%"+desenvolvedora+"%").Find(&jogos)
	c.JSON(200, jogos)
}

func BuscarJogoPorNota(c *gin.Context) {
	var jogos []models.Jogo
	nota := c.Param("nota")
	database.DB.Order("nota desc, titulo").Find(&jogos, "nota >= ?", nota)
	c.JSON(200, jogos)
}
