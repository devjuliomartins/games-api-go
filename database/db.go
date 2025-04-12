package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBancoDeDados() {
	conexao := "host=localhost user=user password=user dbname=games-db port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(conexao))
	if err != nil {
		log.Panic("Erro ao se conectar com o banco de dados")
	}
}
