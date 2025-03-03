package models

import (
	// "time"
	"gorm.io/gorm"
)


// Library represents the Library table
type Library struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
}

// // User represents the Users table
// type User struct {
// 	ID          uint      `gorm:"primary_key;auto_increment" json:"id"`
// 	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
// 	Email       string    `gorm:"type:varchar(255);unique;not null" json:"email"`
// 	ContactNumber string  `gorm:"type:varchar(15)" json:"contact_number"`
// 	Role        string    `gorm:"type:varchar(50)" json:"role"`
// 	LibID       uint      `gorm:"not null" json:"lib_id"`
// 	Library     Library  `gorm:"foreignkey:LibID" json:"library"`
// }

// BookInventory represents the BookInventory table
// type BookInventory struct {
// 	ISBN            string  `gorm:"primary_key;type:varchar(20)" json:"isbn"`
// 	LibID           uint    `gorm:"not null" json:"lib_id"`
// 	Title           string  `gorm:"type:varchar(255);not null" json:"title"`
// 	Authors         string  `gorm:"type:text" json:"authors"`
// 	Publisher       string  `gorm:"type:varchar(255)" json:"publisher"`
// 	Version         string  `gorm:"type:varchar(50)" json:"version"`
// 	TotalCopies     int     `gorm:"not null" json:"total_copies"`
// 	AvailableCopies int     `gorm:"not null" json:"available_copies"`
// 	Library         Library `gorm:"foreignkey:LibID" json:"library"`
// }

// // RequestEvent represents the RequestEvents table
// type RequestEvent struct {
// 	ReqID        uint      `gorm:"primary_key;auto_increment" json:"req_id"`
// 	BookID       string    `gorm:"not null;type:varchar(20)" json:"book_id"`
// 	ReaderID     uint      `gorm:"not null" json:"reader_id"`
// 	RequestDate  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"request_date"`
// 	ApprovalDate time.Time `gorm:"default:null" json:"approval_date"`
// 	ApproverID   uint      `gorm:"not null" json:"approver_id"`
// 	RequestType  string    `gorm:"type:varchar(50)" json:"request_type"`
// 	Book         BookInventory `gorm:"foreignkey:BookID" json:"book"`
// 	Reader       User         `gorm:"foreignkey:ReaderID" json:"reader"`
// 	Approver     User         `gorm:"foreignkey:ApproverID" json:"approver"`
// }

// // IssueRegistry represents the IssueRegistery table
// type IssueRegistry struct {
// 	IssueID          uint      `gorm:"primary_key;auto_increment" json:"issue_id"`
// 	ISBN             string    `gorm:"not null;type:varchar(20)" json:"isbn"`
// 	ReaderID         uint      `gorm:"not null" json:"reader_id"`
// 	IssueApproverID  uint      `gorm:"not null" json:"issue_approver_id"`
// 	IssueStatus      string    `gorm:"type:varchar(50)" json:"issue_status"`
// 	IssueDate        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"issue_date"`
// 	ExpectedReturnDate time.Time `gorm:"not null" json:"expected_return_date"`
// 	ReturnDate       time.Time `gorm:"default:null" json:"return_date"`
// 	ReturnApproverID uint      `gorm:"not null" json:"return_approver_id"`
// 	Book             BookInventory `gorm:"foreignkey:ISBN" json:"book"`
// 	Reader           User         `gorm:"foreignkey:ReaderID" json:"reader"`
// 	IssueApprover    User         `gorm:"foreignkey:IssueApproverID" json:"issue_approver"`
// 	ReturnApprover   User         `gorm:"foreignkey:ReturnApproverID" json:"return_approver"`
// }

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}