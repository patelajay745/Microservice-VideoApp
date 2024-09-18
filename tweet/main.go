package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/patelajay745/Microservice-VideoApp/tweet/config"
	"github.com/patelajay745/Microservice-VideoApp/tweet/routes"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error while loading .env file", err)
	}

	client := config.ConnectDB()

	router := gin.Default()

	router.Use(cors.Default())

	gin.SetMode(os.Getenv("GIN_MODE"))

	routes.SetUpRouter(router, client)

	router.Run(":" + os.Getenv("PORT"))
}
