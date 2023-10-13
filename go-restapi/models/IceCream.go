package models

import "gorm.io/gorm"

type IceCream struct {
	gorm.Model

	Sabor string `gorm:"not null;unique_index" json:"sabor"`
	Tamano string `gorm:"not null" json:"tamano"` 
	Precio float32 `gorm:"not null" json:"precio"`
	Stock  int `gorm:"not null;" json:"stock"`
}