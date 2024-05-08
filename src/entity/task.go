package entity

import "time"

type Task struct {
	ID          int `gorm:"index"`
	UserID      int
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Task) TableName() string {
	return "tasks"
}
