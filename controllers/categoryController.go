package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jayromofo/jayronator/database"
)

func GetCategories(r *database.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Param("category_id")

	}
}

func CreateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	}
}
