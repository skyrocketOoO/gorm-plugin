package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Store struct {
		gorm.Model
		Name     string     `gorm:"unique"`
		ParentID *uuid.UUID `gorm:"type:uuid"`

		Children []*Store `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	}
)
