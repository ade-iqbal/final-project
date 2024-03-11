package model

import "time"

type GormModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt *time.Time 
	UpdatedAt *time.Time 
}