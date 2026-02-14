package main

import (
  "github.com/gin-gonic/gin"
  "github.com/SharathKumarK06/todo-app/internal/database"
  "github.com/SharathKumarK06/todo-app/internal/todo"
)

func main() {
  db, _ := database.Connect()
  // Auto migrate: Create table automatically
  db.AutoMigrate(&todo.Todo{})

  r := gin.Default()

  // Create todo
  r.POST("/todos", func(c *gin.Context) {
    var t todo.Todo

    if err := c.BindJSON(&t); err != nil {
      c.JSON(400, gin.H{"error": "database problem"})
      return
    }

    db.Create(&t)
    c.JSON(201, t)
  })

  // Get list of todos
  r.GET("/todos", func(c *gin.Context) {
    var todos []todo.Todo

    if err := db.Find(&todos).Error; err != nil {
      c.JSON(500, gin.H{"error": err.Error()})
      return
    }

    c.JSON(200, todos)
  })

  // Delete todo
  r.DELETE("/todos/:id", func(c *gin.Context) {
    id := c.Param("id")

    if err := db.Delete(&todo.Todo{}, id).Error; err != nil {
      c.JSON(500, gin.H{"error": "delete failed"})
      return
    }

    c.JSON(200, gin.H{"message": "deleted"})
  })

  // Update todo
  r.PUT("/todos/:id", func(c *gin.Context) {
    id := c.Param("id")

    var existing todo.Todo

    if err := db.First(&existing, id).Error; err != nil {
      c.JSON(404, gin.H{"error": "todo not found"})
      return
    }

    var input todo.Todo
    if err := c.ShouldBindJSON(&input); err != nil {
      c.JSON(400, gin.H{"error": "invalid input"})
      return
    }

    existing.Title = input.Title
    existing.Completed = input.Completed

    db.Save(&existing)

    c.JSON(200, existing)
  })

  r.Run(":8080")
}
