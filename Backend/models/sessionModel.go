// models/session.go
package models

import (
	"time"
	// "gorm.io/gorm"
)

type Session struct {
	ID        uint           `gorm:"primaryKey"`
	UserID    uint           `gorm:"index"`
	Token     string         `gorm:"unique"`
	ExpiresAt time.Time
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}