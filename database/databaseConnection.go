package database

import (
	"github.com/jayromofo/jayronator/models"
	"github.com/jayromofo/jayronator/storage"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
)

type Repository struct {
	DB *gorm.DB
}

func GetClientPGSQL() Repository {
	err := godotenv.Load("./env/.env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("could not load the database")
	}

	r := Repository{
		DB: db,
	}

	return r
}

func MigrateAll(r *Repository) error {
	models.MigrateCategories(r.DB)
	models.MigrateIngredients(r.DB)
	models.MigrateMenu(r.DB)
	models.MigrateRecipe(r.DB)
}
