package models

import "time"

// EmailAttachmentData represents data imported from email attachments (Excel/CSV)
type EmailAttachmentData struct {
	ID           int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	TenantInfoID int       `gorm:"column:tenant_info_id;index" json:"tenant_info_id"`
	FileName     string    `gorm:"column:file_name;type:varchar(255)" json:"file_name"`
	EmailSubject string    `gorm:"column:email_subject;type:varchar(500)" json:"email_subject"`
	EmailFrom    string    `gorm:"column:email_from;type:varchar(255)" json:"email_from"`
	EmailDate    string    `gorm:"column:email_date;type:varchar(100)" json:"email_date"`
	TagData      string    `gorm:"column:tag_data;type:text" json:"tag_data"` // JSON format of all extracted data
	RowData      string    `gorm:"column:row_data;type:longtext" json:"row_data"` // Full row as JSON
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (EmailAttachmentData) TableName() string {
	return "email_attachment_data"
}
