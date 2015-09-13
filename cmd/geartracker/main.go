package main

import (
	"log"
	"os"

	"github.com/smkell/geartracker/Godeps/_workspace/src/github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Static("/", "static")

	router.Run(":" + port)
}
