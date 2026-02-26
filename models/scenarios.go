package models

type App_scenarios struct {
	Appscenarioid int    `gorm:"primaryKey"`
	Scenarioname  string `gorm:"column:scenarioname"`
	Scenarioid    string `gorm:"column:scenarioid"`
	Applicationid string `gorm:"column:applicationid"`
	Userid        int    `gorm:"column:userid"`
	Username      string `gorm:"column:username"`
	Tenantid      int    `gorm:"column:tenantid"`
	Tenantname    string `gorm:"column:tenantname"`
	Status        *int   `gorm:"column:status"`
	// Created      string `gorm:"column:created"`
	// Update       string `gorm:"column:updated"`
}

type Scenariodata struct {
	Appscenarioid int    `gorm:"primaryKey"`
	Scenarioname  string `gorm:"column:scenarioname"`
	Scenarioid    string `gorm:"column:scenarioid"`
	Applicationid string `gorm:"column:applicationid"`
	Userid        int    `gorm:"column:userid"`
	Username      string `gorm:"column:username"`
	Tenantid      string `gorm:"column:tenantid"`
	Tenantname    string `gorm:"column:tenantname"`
	Status        *int   `gorm:"column:status"`
	Created       string `gorm:"column:created"`
	Update        string `gorm:"column:updated"`
}
