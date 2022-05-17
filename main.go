package main

import (
	"day02/model"
	"day02/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
