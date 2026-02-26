package models

import "time"

// EmailDataLogs represents email data imported from various sources
type EmailDataLogs struct {
	ID             int64      `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	TenantInfoID   *int       `gorm:"column:tenant_info_id;index" json:"tenant_info_id"`
	Provider       *string    `gorm:"column:provider;type:varchar(50)" json:"provider"` // gmail, outlook, etc
	EmailFrom      *string    `gorm:"column:email_from;type:varchar(255)" json:"email_from"`
	EmailSubject   *string    `gorm:"column:email_subject;type:varchar(500)" json:"email_subject"`
	EmailBody      *string    `gorm:"column:email_body;type:text" json:"email_body"`
	EmailDate      *string    `gorm:"column:email_date;type:varchar(100)" json:"email_date"`
	SenderName     *string    `gorm:"column:sender_name;type:varchar(255)" json:"sender_name"`
	SenderEmail    *string    `gorm:"column:sender_email;type:varchar(255)" json:"sender_email"`
	RegNo          *string    `gorm:"column:reg_no;type:varchar(50)" json:"reg_no"`
	StockNo        *string    `gorm:"column:stock_no;type:varchar(50)" json:"stock_no"`
	Category       *string    `gorm:"column:category;type:varchar(50)" json:"category"`
	RawData        *string    `gorm:"column:raw_data;type:longtext" json:"raw_data"` // Full JSON of original email
	ProcessedData  *string    `gorm:"column:processed_data;type:text" json:"processed_data"` // Extracted/parsed data
	Status         *int    `gorm:"column:status;type:int;default:0" json:"status"` // pending, processed, failed
	ErrorMessage   *string    `gorm:"column:error_message;type:text" json:"error_message"`
	CreatedAt      *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (EmailDataLogs) TableName() string {
	return "email_data_logs"
}
