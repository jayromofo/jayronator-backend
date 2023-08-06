package models

import (
	"gorm.io/gorm"
	"time"
)

type Menu struct {
	ID         uint       `gorm:"primary key; autoIncrement" json:"id"`
	Name       string     `json:"name"`
	Category   string     `json:"category"`
	Start_Date *time.Time `json:"start_date"`
	End_Date   *time.Time `json:"end_date"`
	Created_at *time.Time `json:"created_at"`
	Updated_at *time.Time `json:"updated_at"`
	Menu_id    string     `json:"menu_id"`
}

func MigrateMenu(db *gorm.DB) error {
	err := db.AutoMigrate(&Menu{})
	return err
}
