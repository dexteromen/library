package models

import (
	"time"
)

// User Table
type User struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Name          string `json:"name" binding:"required"`
	Email         string `gorm:"unique" json:"email" binding:"required,email"`
	ContactNumber string `gorm:"unique" json:"contact_number" binding:"required,len=10,numeric"`
	Password      string `json:"password" binding:"required"`
	Role          string `gorm:"default:'reader'" json:"role,omitempty"`
	LibID         uint   `json:"lib_id,omitempty"`
}

// Library Table
type Library struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"unique" json:"name" binding:"required"`
}

type BookInventory struct {
	BookID uint   `gorm:"primaryKey;autoIncrement" json:"book_id"`
	ISBN   string `gorm:"unique" json:"isbn"`
	// ISBN            string `gorm:"unique;primaryKey;autoIncrement:false;" json:"isbn"`
	LibID           uint   `json:"lib_id,omitempty"`
	Title           string `gorm:"unique" json:"title"`
	Authors         string `json:"authors"`
	Publisher       string `json:"publisher"`
	Version         string `json:"version"`
	TotalCopies     int    `json:"total_copies"`
	AvailableCopies int    `json:"available_copies"`
}

// Request Table
type RequestEvent struct {
	ReqID        uint       `gorm:"primaryKey;autoIncrement" json:"req_id"`
	ISBN         string     `gorm:"not null" json:"isbn" binding:"required"`
	ReaderID     uint       `gorm:"not null" json:"reader_id"`
	RequestDate  string     `gorm:"autoCreateTime" json:"request_date"`
	ApprovalDate *time.Time `json:"approval_date,omitempty"`
	ApproverID   *uint      `json:"approver_id,omitempty"`
	RequestType  string     `gorm:"not null" json:"request_type"`           // e.g., "Borrow", "Return"
	IssueStatus  string     `gorm:"not null" json:"issue_status,omitempty"` // e.g., "Issued and Approved"
}

// Issues Table
type IssueRegistery struct {
	IssueID            uint   `gorm:"primaryKey;autoIncrement" json:"issue_id"`
	ISBN               string `gorm:"not null" json:"isbn"`
	ReaderID           uint   `gorm:"not null" json:"reader_id"`
	IssueApproverID    uint   `gorm:"not null" json:"issue_approver_id"`
	IssueStatus        string `gorm:"not null" json:"issue_status"` // e.g., "Issued", "Returned"
	IssueDate          string `gorm:"autoCreateTime" json:"issue_date"`
	ExpectedReturnDate string `json:"expected_return_date"`
	ReturnDate         string `json:"return_date,omitempty"`
	ReturnApproverID   *uint  `json:"return_approver_id,omitempty"`
}
