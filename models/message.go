package models

import "time"

// This is our single struct/model that is used with Gorm for all
// DB functions.
type Message struct {
  ID        uint      `gorm:"primary_key" json:"id"`
  CreatedAt time.Time                    `json:"-"`
  UpdatedAt time.Time                    `json:"-"`

  Message   string                   `json:"message"`
  UserName  string    `sql:"size:255" json:"user_name"`
  Type      string    `sql:"size:255" json:"type"`
}

type Messages []Message
