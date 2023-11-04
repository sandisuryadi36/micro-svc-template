package db

import "gorm.io/gorm"

type GormProvider struct {
	db_main *gorm.DB
}

func NewProvider(db *gorm.DB) *GormProvider {
	return &GormProvider{db_main: db}
}