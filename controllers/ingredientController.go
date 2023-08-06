package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jayromofo/jayronator/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"time"
)

var ingredientCollection *mongo.Collection = database.OpenCollection(database.Client, "ingredients")
var validate = validator.New()

func GetIngredients() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}

		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}

		startIndex := (page - 1) * recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{"$match", bson.D{{}}}}

		groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"_id", "null"}}},
			{"total_count", bson.D{{"$sum", 1}}},
			{"data", bson.D{{"$push", "$$ROOT"}}}
		}}}

		projectStage := bson.D{
			{
				"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"ingredient_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},

				},
			},
		}
	}
}

func GetIngredient() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateIngredient() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateIngredient() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func DeleteIngredient() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
