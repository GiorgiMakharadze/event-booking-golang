package main

import (
	"github.com/GiorgiMakharadze/event-booking-golang/db"
	"github.com/GiorgiMakharadze/event-booking-golang/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")

}
