package models

import (
	"time"

	"gorm.io/gorm"
)

type Taxpayer struct {
	gorm.Model

	Identifier string `gorm:"uniqueIndex;not null"`
	LegalName  string `gorm:"not null"`

	Receipts []Receipt `gorm:"foreignKey:TaxpayerID"`
}

type Receipt struct {
	gorm.Model

	TaxpayerID         uint
	AccessKey          string    `gorm:"uniqueIndex;not null"`
	Type               string    `gorm:"not null"`
	Establishment      string    `gorm:"not null"`
	EmissionPoint      string    `gorm:"not null"`
	Sequence           string    `gorm:"not null"`
	DateEmission       time.Time `gorm:"not null"`
	DateAuthorization  time.Time `gorm:"not null"`
	ReceptorIdentifier string    `gorm:"not null"`
	ReceptorName       string    `gorm:"not null"`
}
