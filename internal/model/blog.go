package model

import (
	"gorm.io/gorm"
)

type (
	BlogRocord struct {
		gorm.Model

		Children []*Store `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	}
)
