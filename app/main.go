package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Demo struct {
	ID    string `json:"id" gorm:"primarykey"`
	Title string `json:"title"`
}

func main() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, database, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Demo{})

	r := gin.Default()

	r.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		var demo Demo
		res := db.First(&demo, id)
		if res.Error != nil {
			if errors.Is(res.Error, gorm.ErrRecordNotFound) {
				c.JSON(404, gin.H{"error": "Not Found"})
				return
			} else {
				c.JSON(500, gin.H{"error": res.Error.Error()})
			}
		}

		c.JSON(200, demo)
	})

	r.POST("/", func(c *gin.Context) {
		var demo Demo
		err := c.ShouldBindJSON(&demo)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		db.Save(demo)
		c.JSON(204, gin.H{})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"healthy": true,
		})
	})
	r.Run()
}
