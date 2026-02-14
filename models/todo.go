package models

import "time"

type Todo struct {
  ID        uint      `json:"id"          gorm:"primary key"`
  Title     string    `json:"title"       gorm:"not null"       binding:"required,min=3"`
  Completed bool      `json:"completed"   gorm:"default:false"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}
