package main

import (
	"go-blog/model"
	"go-blog/routes"
)

func main() {
	model.InitDB()
	routes.InitRouter()
}