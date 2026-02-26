package models

// VehicleData stores main vehicle info (DVLA + DVSA)
type VehicleData struct {
	ID                       uint        `gorm:"primaryKey"`
	RegistrationNumber       string      `gorm:"uniqueIndex"`
	Make                     string
	Model                    string
	FuelType                 string
	Colour                   string
	MotStatus                string
	MotExpiryDate            string
	TaxStatus                string
	TaxDueDate               string
	YearOfManufacture        int
	EngineCapacity           int
	Co2Emissions             float64
	Wheelplan                string
	MarkedForExport          bool
	MonthOfFirstRegistration string
	FirstUsedDate            string
	RegistrationDate         string
	ManufactureDate          string
	EngineSize               string
	HasOutstandingRecall     string

	MotTests []MotTest `gorm:"foreignKey:VehicleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// MotTest stores MOT test info for a vehicle
type MotTest struct {
	ID                 uint   `gorm:"primaryKey"`
	VehicleID          uint   `gorm:"index"` // foreign key
	CompletedDate      string
	TestResult         string
	ExpiryDate         string
	OdometerValue      string
	OdometerUnit       string
	OdometerResultType string
	MotTestNumber      string
	DataSource         string
	Location           *string `gorm:"null"`

	Defects []MotDefect `gorm:"foreignKey:MotTestID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// MotDefect stores defects for a MOT test
type MotDefect struct {
	ID        uint   `gorm:"primaryKey"`
	MotTestID uint   `gorm:"index"` 
	Text      string
	Type      string
	Dangerous bool
}
