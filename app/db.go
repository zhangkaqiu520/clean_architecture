package app

import "gorm.io/gorm"

var DB *gorm.DB

func NewDB() *gorm.DB {
	return DB
}
