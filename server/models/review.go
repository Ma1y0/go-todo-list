package models

import (
    "time"
)

type Review struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    Score     uint8          `gorm:"not null" json:"score"`
    CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
    Book      Book           `gorm:"foreignKey:BookID" json:"book"`
    BookID    uint           `gorm:"not null" json:"-"`
}

