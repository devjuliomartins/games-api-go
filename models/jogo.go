package models

import "gorm.io/gorm"

type Jogo struct {
	gorm.Model
	Titulo         string  `json:"titulo"`
	Genero         string  `json:"genero"`
	Plataforma     string  `json:"plataforma"`
	AnoLancamento  int     `json:"anoLancamento"`
	Desenvolvedora string  `json:"desenvolvedora"`
	Nota           float64 `json:"nota"`
}
