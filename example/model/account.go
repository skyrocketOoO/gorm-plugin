package model

import (
	"gorm.io/gorm"
)

type (
	Account struct {
		gorm.Model
		UserName    string `gorm:"unique;not null"`
		State       int32  `gorm:"not null;default:1"` // 1: active, 2:inactive
		HashPass    []byte `gorm:"not null"`
		Salt        []byte `gorm:"not null"`
		DisplayName string `gorm:"not null"`
		Phone       Phone  `gorm:"embedded;embeddedPrefix:Phone"`
		Address     `gorm:"embedded;embeddedPrefix:Address"`
		RoleID      uint
		Role        Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	}
)

type Phone struct {
	Number string
}

type Address struct {
	Val string
}
