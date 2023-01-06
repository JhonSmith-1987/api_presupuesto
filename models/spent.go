package models

import "gorm.io/gorm"

type Spent struct {
	gorm.Model

	Detail string `json:"detail"`
	Spent  uint64 `json:"spent"`
}
