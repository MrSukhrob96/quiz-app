package entities

import (
	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          *uuid.UUID  `gorm:"column:id, type:uuid, PRIMARY_KEY, AUTO_INCREMENT" json:"id"`
	Name        string      `gorm:"column:name" json:"name"`
	Branch      interface{} `gorm:"column:branch_id" json:"branch"`
	Office      interface{} `gorm:"column:office_id" json:"office"`
	PhoneNumber string      `gorm:"column:phoneNumber" json:"phoneNumber"`
	Status      string      `gorm:"column:status" json:"status"`
	EntityTimestampms
}

type Users []User
