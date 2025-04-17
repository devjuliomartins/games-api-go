package controller

import (
	"fmt"
	"github.com/devjuliomartins/games-api-go/Paginacao/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"strconv"
)

func PaginarConsulta[T any](c *gin.Context, query *gorm.DB, out *[]T) (model.PaginationMeta, error) {
	pagina, limite, offset, err := getPaginationParams(c)
	if err != nil {
		return model.PaginationMeta{}, err
	}

	var total_jogos int64
	if err := query.Count(&total_jogos).Error; err != nil {
		return model.PaginationMeta{}, err
	}

	if err := query.Offset(offset).Limit(limite).Find(out).Error; err != nil {
		return model.PaginationMeta{}, err
	}

	total_paginas := int(math.Ceil(float64(total_jogos) / float64(limite)))

	meta := model.PaginationMeta{
		TotalJogos:   int(total_jogos),
		Limite:       limite,
		PaginaAtual:  pagina,
		TotalPaginas: total_paginas,
	}

	return meta, nil
}

func getPaginationParams(c *gin.Context) (int, int, int, error) {
	paginaStr := c.DefaultQuery("pagina", "1")
	limiteStr := c.DefaultQuery("limite", "10")

	pagina, err := strconv.Atoi(paginaStr)
	if err != nil || pagina < 1 {
		return 0, 0, 0, fmt.Errorf("parâmetro 'pagina' inválido")
	}

	limite, err := strconv.Atoi(limiteStr)
	if err != nil || limite < 1 {
		return 0, 0, 0, fmt.Errorf("parâmetro 'limite' inválido")
	}

	offset := (pagina - 1) * limite
	return pagina, limite, offset, nil
}
