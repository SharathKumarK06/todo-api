package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/SharathKumarK06/todo-app/internal/database"
  "github.com/SharathKumarK06/todo-app/internal/todo"
)

func main() {
  db, _ := database.Connect()
  // Auto migrate: Create table automatically
  db.AutoMigrate(&todo.Todo{})

  r := gin.Default()

  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "pong"})
  })

  r.POST("/todos", func(c *gin.Context) {
    var t todo.Todo

    if err := c.BindJSON(&t); err != nil {
      c.JSON(400, gin.H{"error": err.Error()})
      return
    }

    db.Create(&t)
    c.JSON(201, t)
  })

  r.GET("/todos", func(c *gin.Context) {
    var todos []todo.Todo

    if err := db.Find(&todos).Error; err != nil {
      c.JSON(500, gin.H{"error": err.Error()})
      return
    }

    c.JSON(200, todos)
  })

  r.DELETE("/todos/:id", func(c *gin.Context) {
    id := c.Param("id")

    if err := db.Delete(&todo.Todo{}, id).Error; err != nil {
      c.JSON(500, gin.H{"error": "delete failed"})
      return
    }

    c.JSON(200, gin.H{"message": "deleted"})
  })

  r.Run(":8080")
}
