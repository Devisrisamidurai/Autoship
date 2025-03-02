package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var db *gorm.DB

func initDB() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=tasks port=5432 sslmode=disable"

	// Check if running tests
	if os.Getenv("TEST_MODE") == "true" {
		dsn = "file::memory:?cache=shared"
	}

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&Task{})
}

func getTasks(c *gin.Context) {
	var tasks []Task
	db.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

func createTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&task)
	c.JSON(http.StatusCreated, task)
}

func main() {
	initDB()
	r := gin.Default()
	r.GET("/tasks", getTasks)
	r.POST("tasks", createTask)
	fmt.Println("server running on port 8000...")
	r.Run(":7000")

}
