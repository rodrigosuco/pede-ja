package main

import (
	"fmt"
	"log"

	"github.com/rodrigosuco/pede-ja/internal/database"
	"github.com/rodrigosuco/pede-ja/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World!")

  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  database.DbConnect()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":3000")
}
