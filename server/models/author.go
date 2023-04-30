package models

import (
    "time"
)

type Author struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    Name      string         `gorm:"not null" json:"name"`
    Bio       string         `gorm:"size:2000" json:"bio"`
    Books     []Book         `gorm:"foreignKey:AuthorID" json:"books"`
    Photo     string         `json:"photo"`
    CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}

