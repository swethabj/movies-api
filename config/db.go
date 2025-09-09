package config

import (
	"fmt"
	"os"
	"time"

	"github.com/swethabj/movies-api/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	user := os.Getenv("DB_USER") // pick from .env file  now refer .env.example file
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", user, pass, host, port, name)
	//user:password@tcp(host:port)/dbname?parseTime=true&loc=Local
	fmt.Println("DSN", dsn)

	// use GORM's logger with minimal output in production (info in dev)
	gormLogger := logger.Default.LogMode(logger.Silent)
	//logger.Silent → no SQL logs (good for production).
	//Can switch to logger.Info in dev to see SQL queries.

	gormConfig := &gorm.Config{
		Logger: gormLogger,
	}

	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		utils.Logger.Fatalf("failed to connect database: %v", err)
	}

	// set connection pool settings (important in industry)
	sqlDB, err := db.DB()
	if err != nil {
		utils.Logger.Fatalf("failed to get db instance: %v", err)
	}
	//When your app talks to MySQL, it needs a TCP connection for each query.
	//Opening a new connection for every single request is slow and expensive.
	//Connection pooling keeps a pool of open connections that can be reused, improving performance.
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Second) // Only affects idle or future connections, not connections currently in use.

	DB = db

	// // Auto migrate tables - good for early-stage projects. For mature apps use migrations.
	// if err := DB.AutoMigrate(&models.Genre{}, &models.Actor{}, &models.Movie{}, &models.MovieActor{}); err != nil {
	// 	utils.Logger.Fatalf("auto migrate failed: %v", err)
	// } //Auto-creates database tables from your Go structs. // Useful for small or early-stage projects. // In production, migrations via tools like golang-migrate are preferred.

	utils.Logger.Info("Database connected and migrated")
}
