package controllers

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/SharathKumarK06/todo-app/config"
  "github.com/SharathKumarK06/todo-app/models"
)

func CreateTodo(c *gin.Context) {
  var todo models.Todo

  if err := c.ShouldBindJSON(&todo); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
    return
  }

  config.DB.Create(&todo)
  c.JSON(http.StatusOK, todo)
}

func GetTodo(c *gin.Context) {
  var todos []models.Todo
  config.DB.Find(&todos)
  c.JSON(http.StatusOK, todos)
}

func DeleteTodo(c *gin.Context) {
  id := c.Param("id")
  config.DB.Delete(&models.Todo{}, id)
  c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}

func UpdateTodo(c *gin.Context) {
  id := c.Param("id")

  var todo models.Todo
  if err := config.DB.First(&todo, id).Error; err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
    return
  }

  var input models.Todo
  if err := c.ShouldBindJSON(&input).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
    return
  }

  todo.Title = input.Title
  todo.Completed = input.Completed

  config.DB.Save(&todo)
  c.JSON(http.StatusOK, todo)
}
