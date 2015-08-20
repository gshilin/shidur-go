package models

import "time"

type Book struct {
  ID        uint      `gorm:"primary_key" json:"id"`
  CreatedAt time.Time                    `json:"-"`
  UpdatedAt time.Time                    `json:"-"`

  Author    string                   `json:"author"`
  Content   string    `sql:"size:255" json:"content"`
  Title     string    `sql:"size:255" json:"title"`
  Slides    string    `sql:"size:255" json:"slides"`
}

type Books []Book
