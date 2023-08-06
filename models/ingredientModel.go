package models

import (
	"gorm.io/gorm"
)

type Ingredient struct {
	ID          uint     `gorm:"primary key;autoIncrement" json:"id"`
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Keywords    []string `json:"keywords"`
}

func MigrateIngredients(db *gorm.DB) error {
	err := db.AutoMigrate(&Ingredient{})
	return err
}
