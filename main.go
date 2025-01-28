package main

import (
	"log"
	"receipt/app/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.SetupRoutes(r)
	log.Println("Server starting on port 8081:")
	err := r.Run(":8081")
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
