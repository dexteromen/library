package models

// User Table
type User struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Name          string `json:"name"`
	Email         string `gorm:"unique" json:"email"`
	ContactNumber string `json:"contact_number"`
	Password      string `json:"-"`
	Role          string `json:"role"` // admin, user, approver
	LibID         uint   `json:"lib_id"`
}