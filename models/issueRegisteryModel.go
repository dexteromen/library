package models

import (
	"time"
)

// IssueRegistery Table
type IssueRegistery struct {
	IssueID            uint       `gorm:"primaryKey;autoIncrement" json:"issue_id"`
	ISBN               string     `gorm:"not null" json:"isbn"`
	ReaderID           uint       `gorm:"not null" json:"reader_id"`
	IssueApproverID    uint       `gorm:"not null" json:"issue_approver_id"`
	IssueStatus        string     `gorm:"not null" json:"issue_status"` // e.g., "Issued", "Returned"
	IssueDate          time.Time  `gorm:"autoCreateTime" json:"issue_date"`
	ExpectedReturnDate string     `json:"expected_return_date"`
	ReturnDate         *time.Time `json:"return_date,omitempty"`
	ReturnApproverID   *uint      `json:"return_approver_id,omitempty"`
}
