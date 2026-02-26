package models

import "time"


type TenantInfo struct {
	TenantInfoID          int        `gorm:"column:tenant_info_id;primaryKey;autoIncrement" json:"tenant_info_id"`

	TenantName            string    `gorm:"column:tenant_name;size:50" json:"tenant_name,omitempty"`
	TenantContactNo       string    `gorm:"column:tenant_contactno;size:50" json:"tenant_contactno,omitempty"`

	AppIntegrationConfigID int       `gorm:"column:app_integration_config_id" json:"app_integration_config_id,omitempty"`

	IntegrationSystem     string    `gorm:"column:integration_system;size:50" json:"integration_system,omitempty"`

	TenantAddress1        string    `gorm:"column:tenant_address1;type:text" json:"tenant_address1,omitempty"`
	TenantEmail           string    `gorm:"column:tenant_email;size:50" json:"tenant_email,omitempty"`
	SupportEmail          string    `gorm:"column:support_email;size:50" json:"support_email,omitempty"`
	TenantSuburb          string    `gorm:"column:tenant_suburb;size:50" json:"tenant_suburb,omitempty"`
	TenantCity            string    `gorm:"column:tenant_city;size:50" json:"tenant_city,omitempty"`
	TenantState           string    `gorm:"column:tenant_state;size:50" json:"tenant_state,omitempty"`

	TenantPostcode        int       `gorm:"column:tenant_postcode;default:0" json:"tenant_postcode,omitempty"`
	TenantAddress2        *string    `gorm:"column:tenant_address2;type:text" json:"tenant_address2,omitempty"`

	ReferenceID           int       `gorm:"column:reference_id;default:0" json:"reference_id,omitempty"`

	TenantToken           string    `gorm:"column:tenant_token;type:text" json:"tenant_token,omitempty"`

	VATRegistrationNo     string    `gorm:"column:vat_registration_no;size:50" json:"vat_registration_no,omitempty"`
	CompanyID             string    `gorm:"column:company_id;size:50" json:"company_id,omitempty"`

	Status                int       `gorm:"column:status;default:1" json:"status,omitempty"`

	CreatedAt             time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt             time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at,omitempty"`
}

// type TenantInfo struct {
// 	// TenantInfoID           int     `gorm:"column:tenant_info_id;primaryKey;autoIncrement" json:"tenant_info_id"`
// 	TenantInfoID int `gorm:"column:tenant_info_id;primaryKey;autoIncrement" json:"tenant_info_id"`
// 	TenantName             string  `gorm:"column:tenant_name;size:50;not null" json:"tenant_name"`
// 	TenantContactNo        string  `gorm:"column:tenant_contactno;size:50;not null" json:"tenant_contactno"`
// 	TenantAddress1          string  `gorm:"column:tenant_address1;size:50;not null" json:"tenant_address1"`
// 	TenantAddress2          string  `gorm:"column:tenant_address2;size:50;not null" json:"tenant_address2"`
// 	TenantSuburb           string  `gorm:"column:tenant_suburb;size:50;not null" json:"tenant_suburb"`
// 	TenantCity             string  `gorm:"column:tenant_city;size:50;not null" json:"tenant_city"`
// 	TenantState            string  `gorm:"column:tenant_state;size:50;not null" json:"tenant_state"`
// 	TenantPostcode         int     `gorm:"column:tenant_postcode;not null;default:0" json:"tenant_postcode"`
// 	ReferenceID            int    `gorm:"column:reference_id" json:"reference_id,omitempty"`
// 	TenantToken            string  `gorm:"column:tenant_token;type:text;not null" json:"tenant_token"`
// 	AppIntegrationConfigID int    `gorm:"column:app_integration_config_id" json:"app_integration_config_id,omitempty"`
// 	IntegrationSystem   string  `gorm:"column:integration_system;size:50;not null" json:"integration_system"`
// 	Status                 int     `gorm:"column:status;not null;default:1" json:"status"`
// 	CreatedAt              time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at,omitempty"`
//     UpdatedAt              time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at,omitempty"`
// 	// CreatedAt               `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
// 	// UpdatedAt              string `gorm:"column:updated_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
// }

func (TenantInfo) TableName() string {
	return "tenant_infos"
}
