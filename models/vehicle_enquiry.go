package models

import "time"

// VehicleEnquiry represents user vehicle enquiries
type VehicleEnquiry struct {
	VehicleEnquiryID int        `gorm:"column:vehicle_enquiry_id;primaryKey;autoIncrement" json:"vehicle_enquiry_id"`
	EnquiryCode      *string    `gorm:"column:enquiry_code;size:50" json:"enquiry_code,omitempty"`
	EnquiryDate      *string `gorm:"column:enquiry_date" json:"enquiry_date,omitempty"`
	CarModel         *string    `gorm:"column:car_model;size:100" json:"car_model,omitempty"`
	EnquiryStatus    *string    `gorm:"column:enquiry_status;size:20" json:"enquiry_status,omitempty"`
	AttendedUser     *string    `gorm:"column:attended_user;size:100" json:"attended_user,omitempty"`
	Notes            *string    `gorm:"column:notes;size:255" json:"notes,omitempty"`
	TenantInfoID     *int       `gorm:"column:tenant_info_id;default:1" json:"tenant_info_id,omitempty"`
	RegNo            *string    `gorm:"column:reg_no;size:50" json:"reg_no,omitempty"`
	StockNo          *string    `gorm:"column:stock_no;size:100" json:"stock_no,omitempty"`
	UserName         *string    `gorm:"column:user_name;size:100" json:"user_name,omitempty"`
	ContactInfo      *string    `gorm:"column:contact_info;size:100" json:"contact_info,omitempty"`
	Email            *string    `gorm:"column:email;size:100" json:"email,omitempty"`
	PostCode         *string    `gorm:"column:post_code;size:100" json:"post_code,omitempty"`
	Category 		*string    `gorm:"column:category;size:100" json:"category,omitempty"`
	EnquiryDetails   *string    `gorm:"column:enquiry_details;size:100" json:"enquiry_details,omitempty"`
	// Responded        *int       `gorm:"column:responded;default:0" json:"responded,omitempty"`
	RespondedVia           *string       `gorm:"column:responded_via;default:1" json:"responded_via,omitempty"`
	RespondedBy			*string    `gorm:"column:responded_by;size:100" json:"responded_by,omitempty"`
	CreatedAt        *time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt        *time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

func (VehicleEnquiry) TableName() string {
	return "vehicle_enquiry"
}
