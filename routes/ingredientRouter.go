package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/jayromofo/jayronator/controllers"
)

func IngredientRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/ingredients", controller.GetIngredients())
	incomingRoutes.GET("/ingredients/:ingredient_id", controller.GetIngredient())
	incomingRoutes.POST("/ingredients", controller.CreateIngredient())
	incomingRoutes.PATCH("/ingredients/:ingredient_id", controller.UpdateIngredient())
	incomingRoutes.DELETE("/ingredients/:ingredient_id", controller.DeleteIngredient())
}
