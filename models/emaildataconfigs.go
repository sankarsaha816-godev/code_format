package models

import "time"

// EmailDataConfigs represents the config table for email import
// Mirrors migrations/011_create_email_data_configs.sql

type EmailDataConfigs struct {
	ID                 int       `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantInfoID       int       `gorm:"column:tenant_info_id" json:"tenant_info_id"`
	Provider           string    `gorm:"column:provider" json:"provider"`
	AccessToken        string    `gorm:"column:access_token" json:"access_token"`
	ClientID           string    `gorm:"column:client_id" json:"client_id"`
	ClientSecret       string    `gorm:"column:client_secret" json:"client_secret"`
	TenantID           string    `gorm:"column:tenant_id" json:"tenant_id"`
	Mailbox            string    `gorm:"column:mailbox" json:"mailbox"`
	MaxResults         int       `gorm:"column:max_results" json:"max_results"`
	Query              string    `gorm:"column:query" json:"query"`
	Since              string    `gorm:"column:since" json:"since"`
	ProcessAttachments bool      `gorm:"column:process_attachments" json:"process_attachments"`
	CreateDynamicTable bool      `gorm:"column:create_dynamic_table" json:"create_dynamic_table"`
	TablePrefix        string    `gorm:"column:table_prefix" json:"table_prefix"`
	CreatedAt          time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at" json:"updated_at"`
}
