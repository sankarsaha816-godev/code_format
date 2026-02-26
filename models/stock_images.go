package models

import "time"

// StockImage represents an image associated with a stock/vehicle
type StockImage struct {
	ID             int64     `gorm:"primaryKey;column:id" json:"id"`
	TenantInfoID   int64     `gorm:"column:tenant_info_id;index" json:"tenant_info_id"`
	StockID        int64     `gorm:"column:stock_id;index" json:"stock_id"`
	StockNo        string    `gorm:"column:stock_no;index" json:"stock_no"`
	Filename       string    `gorm:"column:filename" json:"filename"`
	ObjectKey      string    `gorm:"column:object_key;unique" json:"object_key"`
	PublicURL      string    `gorm:"column:public_url" json:"public_url"`
	FileSize       int64     `gorm:"column:file_size" json:"file_size"`
	ContentType    string    `gorm:"column:content_type;default:image/jpeg" json:"content_type"`
	UploadedAt     time.Time `gorm:"column:uploaded_at;default:CURRENT_TIMESTAMP" json:"uploaded_at"`
	CreatedAt      time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName specifies the table name for StockImage
func (StockImage) TableName() string {
	return "stock_images"
}
