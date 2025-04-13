package main

import (
	"github.com/devjuliomartins/games-api-go/database"
	"github.com/devjuliomartins/games-api-go/routes"
)

func main() {
	database.ConectaBancoDeDados()
	routes.RouteGames()
}
