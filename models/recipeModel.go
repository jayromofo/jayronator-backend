package models

import (
	"gorm.io/gorm"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recipe struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `json:"name"	validate:"required"`
	Description string             `json:"description" validate:"required"`
	ImgURL      string             `json:"img_url"`
	Created_at  *time.Time         `json:"created_at"`
	Updated_at  *time.Time         `json:"updated_at"`
	Recipe_id   string             `json:"recipe_id"`
}

func MigrateRecipe(db *gorm.DB) error {
	err := db.AutoMigrate(&Recipe{})
	return err
}
