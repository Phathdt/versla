package sdkcm

import (
	"time"
)

type Model struct {
	ID        uint32    `json:"id" gorm:"id,PRIMARY_KEY"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
