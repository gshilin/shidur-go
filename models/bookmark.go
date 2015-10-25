package models

import "time"

type Bookmark struct {
  ID        uint      `gorm:"primary_key" json:"id"`
  CreatedAt time.Time                    `json:"-"`
  UpdatedAt time.Time                    `json:"-"`

  Author    string                   `json:"author"`
  Book      string    `sql:"size:255" json:"book"`
  BookName  string    `sql:"size:255" json:"book_name"`
  Page      uint      `sql:"size:255" json:"page"`
  Letter    string    `sql:"size:255" json:"letter"`
  Position  uint      `sql:"size:255" json:"position"`
}

type Bookmarks []Bookmark
