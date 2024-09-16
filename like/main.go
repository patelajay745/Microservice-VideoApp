package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/patelajay745/Microservice-VideoApp/like/config"
	"github.com/patelajay745/Microservice-VideoApp/like/routes"
)

func main() {

	// loads .env in root directory
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//connect database
	client := config.ConnectDB()

	router := gin.Default()

	//router.Use(gin.Logger())
	router.Use(cors.Default())

	routes.SetupRouter(router, client)

	router.Run(":8001")

}
