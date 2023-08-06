package models

import (
	"gorm.io/gorm"
)

type Category struct {
	ID          uint   `gorm:"primary key; autoIncrement" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func MigrateCategories(db *gorm.DB) error {
	err := db.AutoMigrate(&Category{})
	return err
}
