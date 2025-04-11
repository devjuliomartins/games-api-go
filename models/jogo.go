package models

type Jogo struct {
	Id             int    `json:"id"`
	Titulo         string `json:"titulo"`
	Genero         string `json:"genero"`
	AnoLancamento  string `json:"anoLancamento"`
	Desenvolvedora string `json:"desenvolvedora"`
	Nota           string `json:"nota"`
}
