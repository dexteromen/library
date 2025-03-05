package models

import (
	"time"
)

// RequestEvent Table
type RequestEvent struct {
	ReqID        uint       `gorm:"primaryKey;autoIncrement" json:"req_id"`
	BookID       uint       `gorm:"not null" json:"book_id"`
	ReaderID     uint       `gorm:"not null" json:"reader_id"`
	RequestDate  time.Time  `gorm:"autoCreateTime" json:"request_date"`
	ApprovalDate *time.Time `json:"approval_date,omitempty"`
	ApproverID   *uint      `json:"approver_id,omitempty"`
	RequestType  string     `gorm:"not null" json:"request_type"` // e.g., "Borrow", "Return"
}
