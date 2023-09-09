package models

import "gorm.io/gorm"

type Worker struct {
	gorm.Model
	Name     string `json:"name"`
	Position string `json:"position"`
	IsAdmin  bool   `json:"isadmin"`
}
