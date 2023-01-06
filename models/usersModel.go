package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model

	Name     string `gorm:"not null" json:"name"`
	User     string `gorm:"not null;unique_index" json:"user"`
	Password string `gorm:"not null" json:"password"`
}
