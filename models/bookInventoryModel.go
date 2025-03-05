package models

// import (
// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// BookInventory Table
type BookInventory struct {
	// BookID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"book_id"`
	BookID          uint   `gorm:"primaryKey;autoIncrement" json:"book_id"`
	ISBN            string `gorm:"primaryKey;autoIncrement:false" json:"isbn"`
	LibID           uint   `json:"lib_id"`
	Title           string `json:"title"`
	Authors         string `json:"authors"`
	Publisher       string `json:"publisher"`
	Version         string `json:"version"`
	TotalCopies     int    `json:"total_copies"`
	AvailableCopies int    `json:"available_copies"`
}

// // BeforeCreate ensures a UUID is assigned before inserting into the database
// func (b *BookInventory) BeforeCreate(tx *gorm.DB) (err error) {
// 	if b.BookID == uuid.Nil {
// 		b.BookID = uuid.New()
// 	}
// 	return
// }
