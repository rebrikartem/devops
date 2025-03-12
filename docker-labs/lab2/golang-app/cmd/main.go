package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"unique"`
}

var db *gorm.DB

func main() {
	// Инициализация базы данных
	dsn := os.Getenv("DATABASE_URL")
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Запуск миграций (если требуется)
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		runMigrations()
		return
	}

	// Запуск веб-сервера
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" // Значение по умолчанию
	}
	log.Fatal(r.Run(":" + port))
}

func runMigrations() {
	m, err := migrate.New(
		"file:///app/migrations",
		os.Getenv("DATABASE_URL"),
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Migrations applied successfully")
}
