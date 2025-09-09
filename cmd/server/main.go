package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/swethabj/movies-api/config"
	_ "github.com/swethabj/movies-api/docs"
	"github.com/swethabj/movies-api/internal/server"
	"github.com/swethabj/movies-api/utils"
)

// @title Movies API
// @version 1.0
// @description This is a sample movies API with Fiber and Swagger.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// load .env
	_ = godotenv.Load("/Users/Swetha.BJ/Desktop/Personal/Project/Go/CRUD_fiber_Mysql/movies-api/.env")
	fmt.Println("Loaded Env File")

	utils.InitLogger(os.Getenv("LOG_LEVEL"))

	// connect db
	config.ConnectDB()

	app := server.NewRouter()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	app.Use(cors.New())

	utils.Logger.Infof("starting server on :%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
