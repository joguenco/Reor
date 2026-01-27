package models

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model

	Identifier     string    `gorm:"uniqueIndex;not null"`
	Name           string    `gorm:"not null"`
	FriendlyName   string    `gorm:"not null"`
	Email          string    `gorm:"not null"`
	Token          string    `gorm:"not null"`
	Status         string    `gorm:"not null"`
	LimitedBy      string    `gorm:"not null"`
	Limit          uint      `gorm:"not null"`
	ExpirationDate time.Time `gorm:"not null"`
	Observations   *string
	RoleID         uint
}

type Role struct {
	gorm.Model

	Name        string `gorm:"uniqueIndex;not null"`
	Description string `gorm:"not null"`

	Subscriptions []Subscription `gorm:"foreignKey:RoleID"`
}
