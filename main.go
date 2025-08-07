package main

import (
	"eventsManagement/db"
	"eventsManagement/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	//Initialize database and Gin Server
	db.InitDBConnection()
	server := gin.Default()

	//Regester routes for all requests
	routes.RegisterRoutes(server)

	//Run the server
	server.Run(":8080")

}
