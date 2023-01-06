package models

import "gorm.io/gorm"

type Budget struct {
	gorm.Model

	Detail string `json:"detail"`
	Budget uint64 `json:"budget"`
}
