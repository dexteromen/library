package models

// // User Table
// type User struct {
// 	ID            uint   `gorm:"primaryKey" json:"id"`
// 	Name          string `json:"name"`
// 	Email         string `gorm:"unique" json:"email"`
// 	ContactNumber string `json:"contact_number"`
// 	Password      string `json:"-"`
// 	Role          string `json:"role"` // admin, user, approver
// 	LibID         uint   `json:"lib_id"`
// }

// User Table
type User struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Name          string `json:"name" binding:"required"`
	Email         string `gorm:"unique" json:"email" binding:"required,email"`
	ContactNumber string `json:"contact_number" binding:"required,len=10,numeric"`
	Password      string `json:"password" binding:"required"`
	Role          string `json:"role" binding:"required,oneof=admin reader owner"`
	LibID         uint   `json:"lib_id"`
}
