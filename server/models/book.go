package models

import (
    "time"
)

type Book struct {
    ID          uint           `gorm:"primaryKey" json:"id"`
    Name        string         `gorm:"not null" json:"name"`
    Author      Author         `gorm:"foreignKey:AuthorID" json:"author"`
    AuthorID    uint           `gorm:"not null" json:"-"`
    CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
    Reviews     []Review       `gorm:"foreignKey:BookID" json:"reviews"`
    Description string         `gorm:"size:1000" json:"description"`
    Cover       string         `json:"cover"`
}

