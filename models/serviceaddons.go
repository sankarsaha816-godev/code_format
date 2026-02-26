package models

import "time"

type ServiceType struct {
	ServiceTypeID   int     `gorm:"column:service_type_id;primaryKey;autoIncrement" json:"service_type_id"`
	ServiceCategory string  `gorm:"column:service_category;type:varchar(100);not null" json:"service_category"`
	ServiceType     string  `gorm:"column:service_type;type:varchar(100);not null" json:"service_type"`
	TenantInfoID  int     `gorm:"column:tenant_info_id" json:"tenant_info_id"`
	Status          int     `gorm:"column:status;type:int;default:1;not null" json:"status"`
	CreatedBy       *string `gorm:"column:created_by;type:varchar(50)" json:"created_by"`
	UpdatedBy       *string `gorm:"column:updated_by;type:varchar(50)" json:"updated_by"`
	// CreatedAt       *string `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	// UpdatedAt       *string `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP" json:"updated_at"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	// Relation: One ServiceType has many ServiceDetails
	ServiceDetails []ServiceDetail `gorm:"foreignKey:ServiceTypeID;references:ServiceTypeID" json:"service_details"`
}

type ServiceDetail struct {
	ServiceDetailID int     `gorm:"column:service_detail_id;primaryKey;autoIncrement" json:"service_detail_id"`
	ServiceTypeID   *int    `gorm:"column:service_type_id" json:"service_type_id"`
	ServiceType     *string `gorm:"column:service_type;type:varchar(50)" json:"service_type"`
	ServiceDetail   *string `gorm:"column:service_detail;type:varchar(50)" json:"service_detail"`
	Status          int     `gorm:"column:status;type:int;default:1;not null" json:"status"`
	CreatedBy       *string `gorm:"column:created_by;type:varchar(50)" json:"created_by"`
	UpdatedBy       *string `gorm:"column:updated_by;type:varchar(50)" json:"updated_by"`
	CreatedAt       time.Time    `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time    `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	// CreatedAt       *string `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	// UpdatedAt       *string `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP" json:"updated_at"`
}

// type VasSales struct {
// 	VasSalesID              int        `gorm:"primaryKey;autoIncrement;column:vas_sales_id"`
// 	TenantInfoID			   int        `gorm:"column:tenant_info_id"`
// 	ServiceDetailID             int        `gorm:"column:service_detail_id"`
// 	StockNumber          string     `gorm:"column:stock_number;size:50"`
// 	CustomerID                  int        `gorm:"column:customer_id"`
// 	SalespersonID               int        `gorm:"column:salesperson_id"`
// 	Provider                    string     `gorm:"column:provider;size:100"`
// 	StartDate                   *time.Time `gorm:"column:start_date"`
// 	EndDate                     *time.Time `gorm:"column:end_date"`
// 	Status                      int     `gorm:"column:status"`
// 	NetPriceExVAT               float64    `gorm:"column:net_price_ex_vat"`
// 	RetailPriceIncVAT           float64    `gorm:"column:retail_price_inc_vat"`
// 	VATRate                     float64    `gorm:"column:vat_rate"`
// 	VATReceipt                  string     `gorm:"column:vat_receipt;size:100"`
// 	VATAmount                   float64    `gorm:"column:vat_amount"`
// 	CostPrice                   float64    `gorm:"column:cost_price"`
// 	RefundedAmount			  float64    `gorm:"column:refunded_amount"`
// 	FCADisclosureStatus         string     `gorm:"column:fca_disclosure_status;size:50"`
// 	CustomerAcceptanceDate      *time.Time `gorm:"column:customer_acceptance_date"`
// 	CancellationRightsExplained *bool      `gorm:"column:cancellation_rights_explained"`
// 	InvoiceNumber               string     `gorm:"column:invoice_number;size:100"`
// 	DeliveryMethod              string     `gorm:"column:delivery_method;size:50"`
// 	ActivationCode              string     `gorm:"column:activation_code;size:100"`
// 	FulfillmentStatus           string     `gorm:"column:fulfillment_status;size:50"`
// 	Notes                       string     `gorm:"column:notes;type:text"`
// 	CreatedAt                   *time.Time `gorm:"column:created_at;autoCreateTime"`
// 	UpdatedAt                   *time.Time `gorm:"column:updated_at;autoUpdateTime"`
// }

	type VasSales struct {
		VasSalesID                int        `gorm:"primaryKey;autoIncrement;column:vas_sales_id"`                                                          // int(10) NOT NULL AUTO_INCREMENT
		ServiceTypeID   *int     `gorm:"column:service_type_id"`
		ServiceCategory *string  `gorm:"column:service_category"`
		ServiceType     *string  `gorm:"column:service_type"`
		RegNO 				   *string    `gorm:"column:reg_no"`
		                                                                                          // varchar(50) NULL
		ServiceDetailID           *int       `gorm:"column:service_detail_id"` 
		ServiceDetail *string  `gorm:"column:service_detail"`                                                                           // int(10) NULL
		TenantInfoID              *int       `gorm:"column:tenant_info_id"`                                                                                  // int(10) NULL
		StockNumber               *string    `gorm:"column:stock_number;size:50"`                                                                            // varchar(50) NULL
		CustomerID                *int       `gorm:"column:customer_id"`
		// CustomerID *int	`gorm:"column:customer_id" json:"customer_id"`
		CustomerName *string `gorm:"column:customer_name;size:100" json:"customer_name"`                                                                                     // int(10) NULL
		SalespersonID             *int       `gorm:"column:salesperson_id"`                                                                                  // int(10) NULL
		Provider                  *string    `gorm:"column:provider;size:100"`                                                                               // varchar(100) NULL
		StartDate                 *string `gorm:"column:start_date"`                                                                                      // date NULL
		EndDate                   *string `gorm:"column:end_date"`                                                                                        // date NULL
		Status                    *int       `gorm:"column:status;default:1"`                                                                                // int(10) NULL default 1 (1-Active,2-Expired,3-Cancelled,4-Pending)
		NetPriceExVAT             *float64   `gorm:"column:net_price_ex_vat"`                                                                               // decimal(10,2) NULL
		RetailPriceIncVAT         *float64   `gorm:"column:retail_price_inc_vat"`                                                                           // decimal(10,2) NULL
		VATRate                   *float64   `gorm:"column:vat_rate"`                                                                                       // decimal(5,2) NULL
		VATReceipt                *string    `gorm:"column:vat_receipt;size:100"`                                                                            // varchar(100) NULL
		VATAmount                 *float64   `gorm:"column:vat_amount"`                                                                                      // decimal(10,2) NULL
		CostPrice                 *float64   `gorm:"column:cost_price"` // decimal(10,2) NULL
		SaleDate 				*string `gorm:"column:sale_date"`                                                                                                                                                                       
		RefundedAmount            *float64   `gorm:"column:refunded_amount"`                                                                                // decimal(10,2) NULL
		FCADisclosureStatus       *string    `gorm:"column:fca_disclosure_status;size:50"`                                                                   // varchar(50) NULL
		CustomerAcceptanceDate    *string `gorm:"column:customer_acceptance_date"`                                                                        // datetime NULL
		CancellationRightsExplained *bool    `gorm:"column:cancellation_rights_explained"`                                                                   // tinyint(1) NULL
		InvoiceNo             *string    `gorm:"column:invoice_no;size:100"`                                                                         // varchar(100) NULL
		DeliveryMethod            *string    `gorm:"column:delivery_method;size:50"`                                                                         // varchar(50) NULL
		ActivationCode            *string    `gorm:"column:activation_code;size:100"`                                                                        // varchar(100) NULL
		FulfillmentStatus         *string    `gorm:"column:fulfillment_status;size:50"`                                                                      // varchar(50) NULL
		Notes                     *string    `gorm:"column:notes;type:text"`                                                                                 // text NULL
		CreatedAt                 *time.Time `gorm:"column:created_at;autoCreateTime"`                                                                       // datetime NULL DEFAULT CURRENT_TIMESTAMP
		UpdatedAt                 *time.Time `gorm:"column:updated_at;autoUpdateTime"`                                                                       // datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	}
