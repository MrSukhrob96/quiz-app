package entities

import (
	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID     *uuid.UUID `gorm:"column:id, type:uuid, PRIMARY_KEY, AUTO_INCREMENT" json:"id"`
	Name   string     `gorm:"column:name" json:"name"`
	Status string     `gorm:"column:status" json:"status"`
	EntitySoftDeletes
}

type Roles []Role
