package main

import (
	"github.com/guilhermeonrails/api-go-gin/database"
	"github.com/guilhermeonrails/api-go-gin/routes"
)

//testando comentario e pull request actions go

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
