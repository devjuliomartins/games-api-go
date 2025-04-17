package model

type PaginationMeta struct {
	TotalJogos   int `json:"total_jogos"`
	Limite       int `json:"limite"`
	PaginaAtual  int `json:"pagina_atual"`
	TotalPaginas int `json:"total_paginas"`
}
