package entities

import "time"

type EntityTimestampms struct {
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type EntitySoftDeletes struct {
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
