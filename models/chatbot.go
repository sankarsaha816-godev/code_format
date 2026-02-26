package models

import "gorm.io/datatypes"

type KnowledgeBases struct {
	ID           uint      `gorm:"primaryKey"`
	TenantInfoID int      `gorm:"not null"`
	Question     string    `gorm:"unique;not null"`
	Answer       string    `gorm:"not null"`
	Embedding    datatypes.JSON `gorm:"type:json"`
}

type ChatLog struct {
	ID           uint    `gorm:"primaryKey"`
	TenantInfoID uint    `gorm:"not null"`
	Email        string  `gorm:"not null"`
	Question     string  `gorm:"not null"`
	Answer       string  `gorm:"not null"`
	Feedback     string  `gorm:"type:text;default:'good'"`
	Learned      bool    `gorm:"default:false"`
	Confidence   float64 `gorm:"default:0"`
}
