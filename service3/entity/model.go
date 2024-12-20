package entity

import (
	"time"
)

type Model struct {
	ID        int64 `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
}

func New(name string) Model {
	return Model{
		Name:      name,
		CreatedAt: time.Now(),
	}
}