package models

import (
	"time"
)

type Question struct {
	ID        int       `gorm:"primaryKey NOT NULL index" column:"id" json:"id"`
	CreatedAt time.Time `gorm:"column:created" json:"created"`
	UpdatedAt time.Time `gorm:"column:updated" json:"updated"`
}
