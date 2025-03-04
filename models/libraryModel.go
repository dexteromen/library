package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Library Table
type Library struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

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

// BookInventory Table
type BookInventory struct {
	BookID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"book_id"`
	ISBN            string    `gorm:"primaryKey" json:"isbn"`
	LibID           uint      `json:"lib_id"`
	Title           string    `json:"title"`
	Authors         string    `json:"authors"`
	Publisher       string    `json:"publisher"`
	Version         string    `json:"version"`
	TotalCopies     int       `json:"total_copies"`
	AvailableCopies int       `json:"available_copies"`
}

// BeforeCreate ensures a UUID is assigned before inserting into the database
func (b *BookInventory) BeforeCreate(tx *gorm.DB) (err error) {
	if b.BookID == uuid.Nil {
		b.BookID = uuid.New()
	}
	return
}

// RequestEvents Table
type RequestEvents struct {
	ReqID        uint      `gorm:"primaryKey" json:"req_id"`
	BookID       string    `json:"book_id"`
	ReaderID     uint      `json:"reader_id"`
	RequestDate  time.Time `json:"request_date"`
	ApprovalDate time.Time `json:"approval_date"`
	ApproverID   uint      `json:"approver_id"`
	RequestType  string    `json:"request_type"` // borrow, return
}

// IssueRegistery Table
type IssueRegistery struct {
	IssueID            uint      `gorm:"primaryKey" json:"issue_id"`
	ISBN               string    `json:"isbn"`
	ReaderID           uint      `json:"reader_id"`
	IssueApproverID    uint      `json:"issue_approver_id"`
	IssueStatus        string    `json:"issue_status"` // issued, returned
	IssueDate          time.Time `json:"issue_date"`
	ExpectedReturnDate time.Time `json:"expected_return_date"`
	ReturnDate         time.Time `json:"return_date"`
	ReturnApproverID   uint      `json:"return_approver_id"`
}
