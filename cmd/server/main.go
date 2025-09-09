package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/swethabj/movies-api/config"
	"github.com/swethabj/movies-api/internal/server"
	"github.com/swethabj/movies-api/utils"
)

func main() {
	// load .env
	_ = godotenv.Load("../../.env")
	fmt.Println("Loaded Env File")

	utils.InitLogger(os.Getenv("LOG_LEVEL"))

	// connect db
	config.ConnectDB()

	app := server.NewRouter()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	utils.Logger.Infof("starting server on :%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
