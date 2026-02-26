package models

import "time"

type SalesInvoiceHeader struct {
    SalesInvoiceHeaderID int        `gorm:"column:sales_invoice_header_id;primaryKey;autoIncrement" json:"sales_invoice_header_id"`
    InvoiceNumber        *string    `gorm:"column:invoice_number" json:"invoice_number"`
    InvoiceDate          *string    `gorm:"column:invoice_date" json:"invoice_date"`
    CustomerID           *int       `gorm:"column:customer_id" json:"customer_id"`
    CustomerName         *string    `gorm:"column:customer_name" json:"customer_name"`
    StockNumber          *string    `gorm:"column:stock_number" json:"stock_number"`
    RegNo                *string    `gorm:"column:reg_no" json:"reg_no"`
    TenantInfoID         *int       `gorm:"column:tenant_info_id" json:"tenant_info_id"`
    TenantName           *string    `gorm:"column:tenant_name" json:"tenant_name"`
    Status               int        `gorm:"column:status;default:1" json:"status"`
    CreatedBy            *string    `gorm:"column:created_by" json:"created_by"`
    UpdatedBy            *string    `gorm:"column:updated_by" json:"updated_by"`
    CreatedAt            *time.Time `gorm:"column:created_at" json:"created_at"`
    UpdatedAt            *time.Time `gorm:"column:updated_at" json:"updated_at"`
    Email        *string `gorm:"column:email" json:"email"`
    PaidAmount *float64 `gorm:"column:paid_amount" json:"paid_amount"`
    AddressLine1 *string `gorm:"column:address_line1" json:"address_line1"`
    AddressLine2 *string `gorm:"column:address_line2" json:"address_line2"`
    AddressLine3 *string `gorm:"column:address_line3" json:"address_line3"`
    City         *string `gorm:"column:city" json:"city"`
    State        *string `gorm:"column:state" json:"state"`
    Country      *string `gorm:"column:country" json:"country"`
    PostCode     *string `gorm:"column:post_code" json:"post_code"`
	Details []SalesInvoiceDetail `gorm:"foreignKey:SalesInvoiceHeaderID;references:SalesInvoiceHeaderID" json:"details"`
}


// type SalesInvoiceHeader struct {
// 	SalesInvoiceHeaderID int        `gorm:"column:sales_invoice_header_id;primaryKey;autoIncrement" json:"sales_invoice_header_id"`
// 	InvoiceNumber        *string    `gorm:"column:invoice_number" json:"invoice_number"`
// 	InvoiceDate          *string    `gorm:"column:invoice_date" json:"invoice_date"`
// 	CustomerID           *int       `gorm:"column:customer_id" json:"customer_id"`
// 	CustomerName         *string    `gorm:"column:customer_name" json:"customer_name"`
// 	StockNumber          *string    `gorm:"column:stock_number" json:"stock_number"`
// 	RegNo                *string    `gorm:"column:reg_no" json:"reg_no"`
// 	TenantInfoID         *int       `gorm:"column:tenant_info_id" json:"tenant_info_id"`
// 	TenantName           *string    `gorm:"column:tenant_name" json:"tenant_name"`
// 	Status               int        `gorm:"column:status;default:1" json:"status"`
// 	CreatedBy            *string    `gorm:"column:created_by" json:"created_by"`
// 	UpdatedBy            *string    `gorm:"column:updated_by" json:"updated_by"`
// 	CreatedAt            *time.Time `gorm:"column:created_at" json:"created_at"`
// 	UpdatedAt            *time.Time `gorm:"column:updated_at" json:"updated_at"`
// 	// Child relationship (One-to-Many)
// 	Details []SalesInvoiceDetail `gorm:"foreignKey:SalesInvoiceHeaderID;references:SalesInvoiceHeaderID" json:"details"`
// }

type SalesInvoiceDetail struct {
	SalesInvoiceDetailID int        `gorm:"column:sales_invoice_detail_id;primaryKey;autoIncrement" json:"sales_invoice_detail_id"`
	SalesInvoiceHeaderID *int       `gorm:"column:sales_invoice_header_id" json:"sales_invoice_header_id"`
	InvoiceNumber        *string    `gorm:"column:invoice_number" json:"invoice_number"`
	RegNo                *string    `gorm:"column:reg_no" json:"reg_no"`
	Description          *string    `gorm:"column:description" json:"description"`
	StartDate            *string    `gorm:"column:start_date" json:"start_date"`
	EndDate              *string    `gorm:"column:end_date" json:"end_date"`
	NetPrice             *float64   `gorm:"column:net_price" json:"net_price"`
	Vat                  *string    `gorm:"column:vat" json:"vat"`
	RetailPrice          *float64   `gorm:"column:retail_price" json:"retail_price"`
	Quantity             *int       `gorm:"column:quantity" json:"quantity"`
	Status               int        `gorm:"column:status;default:1" json:"status"`
    SalesID             *int       `gorm:"column:sales_id" json:"sales_id"`
    VasSalesID          *int       `gorm:"column:vas_sales_id" json:"vas_sales_id"`
    Category           *string    `gorm:"column:category" json:"category"`
	CreatedBy            *string    `gorm:"column:created_by" json:"created_by"`
	UpdatedBy            *string    `gorm:"column:updated_by" json:"updated_by"`
	CreatedAt            *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt            *time.Time `gorm:"column:updated_at" json:"updated_at"`
}
