package models

import (
	"time"
)

type BaseModel struct {
	UUID string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
