package main

import (
	"fmt"
	"log"
	"net/http"

	_ "gowebgame/docs"

	Cache "gowebgame/internal/redis"
	Database "gowebgame/internal/db"
	Router "gowebgame/internal/router"
)

// @title Gowebgame API
// @version 1.0
// @description API de base pour jeu en Go
// @host localhost:8181
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" suivi d'un espace et de votre token JWT.
func main() {
	Cache.InitRedis()
	Database.InitSQLite()

	r := Router.InitRoute()

	fmt.Println("Server Started on :8181")
	log.Fatal(http.ListenAndServe(":8181", r))
}
