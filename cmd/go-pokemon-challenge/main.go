package main

import (
	"github.com/mikaellpc4/go-pokemon-challenge/initializers"
	"github.com/mikaellpc4/go-pokemon-challenge/internal/routes"
)

func init() {
	initializers.LoadEnv()
	initializers.InitializeDB()
}

func main() {
	router := routes.NewRouter()

	port := "3333"
	err := router.Start(":" + port)

	if err != nil {
		router.Logger.Fatal(err)
	}
}
