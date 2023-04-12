package entities

import (
	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Office struct {
	gorm.Model
	ID     uuid.UUID `gorm:"column:id, type:uuid, PRIMARY_KEY, AUTO_INCREMENT" json:"id"`
	Name   string    `gorm:"column:name" json:"name"`
	Branch Branch    `gorm:"foreignKey:ID"`
	Status string    `gorm:"column:status" json:"status"`
}
