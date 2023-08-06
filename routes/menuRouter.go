package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/jayromofo/jayronator/controllers"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menu", controller.GetIngredients())
	incomingRoutes.GET("/menu/:menu_id", controller.GetIngredient())
	incomingRoutes.POST("/menu", controller.CreateIngredient())
	incomingRoutes.PATCH("/menu/:menu_id", controller.UpdateIngredient())
	incomingRoutes.DELETE("/menu/:menu_id", controller.DeleteIngredient())
}
