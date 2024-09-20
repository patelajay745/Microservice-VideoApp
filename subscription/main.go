package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/patelajay745/Microservice-VideoApp/subscription/config"
	"github.com/patelajay745/Microservice-VideoApp/subscription/routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	router := gin.Default()

	router.Use(cors.Default())

	routes.SetUpRouter(router)

	router.Run(":" + os.Getenv("PORT"))

}
