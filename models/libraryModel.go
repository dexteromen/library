package models

// Library Table
type Library struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
