package models

import "time"

type CustomerInfo struct {
    CustomerInfoID int        `gorm:"column:customer_info_id;primaryKey;autoIncrement" json:"customer_info_id"`
    CustomerName   string     `gorm:"column:customer_name;size:150;not null" json:"customer_name"`
    FirstName      string     `gorm:"column:first_name;size:150;not null" json:"first_name"`
    LastName       string     `gorm:"column:last_name;size:150;not null" json:"last_name"`
    EmailID        string     `gorm:"column:email_id;size:150;not null" json:"email_id"`
    ContactNo1     string     `gorm:"column:contact_no1;size:100;not null" json:"contact_no1"`
    ContactNo2     string     `gorm:"column:contact_no2;size:100;not null" json:"contact_no2"`

    TenantInfoID   int        `gorm:"column:tenant_info_id;not null" json:"tenant_info_id"`
    TenantName     string     `gorm:"column:tenant_name;size:100;not null" json:"tenant_name"`

    AddressLine1   string     `gorm:"column:address_line1;type:text;not null" json:"address_line1"`
    AddressLine2   string     `gorm:"column:address_line2;type:text;not null" json:"address_line2"`
    AddressLine3   string     `gorm:"column:address_line3;type:text;not null" json:"address_line3"`

    City           string     `gorm:"column:city;size:100;not null" json:"city"`
    State          string     `gorm:"column:state;size:100;not null" json:"state"`
    Country        string     `gorm:"column:country;size:100;not null" json:"country"`
    PostCode       string     `gorm:"column:post_code;size:50;not null" json:"post_code"`

    Status         int        `gorm:"column:status;default:1" json:"status"`
    CreatedBy      string     `gorm:"column:created_by;size:50;not null" json:"created_by"`
    UpdatedBy      string     `gorm:"column:updated_by;size:50;not null" json:"updated_by"`

    CreatedAt      *time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt      *time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
}

func (CustomerInfo) TableName() string {
    return "customer_info"
}
