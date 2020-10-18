package sdkcm

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint32         `json:"id" gorm:"id,PRIMARY_KEY"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
