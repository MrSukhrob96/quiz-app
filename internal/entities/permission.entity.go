package entities

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	ID     *uuid.UUID `gorm:"column:id, type:uuid, PRIMARY_KEY, AUTO_INCREMENT" json:"id"`
	Name   string     `gorm:"column:name" json:"name"`
	Status string     `gorm:"column:status" json:"status"`
	EntityTimestampms
}

type Permissions []Permission
