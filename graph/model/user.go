package model

import "github.com/google/uuid"

type User struct {
	ID   uuid.UUID `gorm:"type:varchar(128); primaryKey"`
	Name *string
}
