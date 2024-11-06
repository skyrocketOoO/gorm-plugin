package model

import (
	"gorm.io/gorm"
)

// Any schema change must make sure master role has all permissions

type (
	Role struct {
		gorm.Model
		Name string `gorm:"not null;uniqueIndex:idx_store_name"`
		Page Page   `gorm:"embedded; embeddedPrefix:Page"`

		Accounts []Account
	}

	PageOperation struct {
		Create bool `gorm:"not null;default:false"`
		Read   bool `gorm:"not null;default:false"`
		Update bool `gorm:"not null;default:false"`
		Delete bool `gorm:"not null;default:false"`
	}

	Page struct {
		// LoginRecord PageOperation `gorm:"embedded;embeddedPrefix:LoginRecord"`
		Role    PageOperation `gorm:"embedded;embeddedPrefix:Role"`
		Setting PageOperation `gorm:"embedded;embeddedPrefix:Setting"`
		Account PageOperation `gorm:"embedded;embeddedPrefix:Account"`
	}
)
