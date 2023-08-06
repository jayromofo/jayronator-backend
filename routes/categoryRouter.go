package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/jayromofo/jayronator/controllers"
)

func CategoryRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/category", controller.GetCategories())
	incomingRoutes.GET("/category/:category_id", controller.GetIngredient())
	incomingRoutes.POST("/category", controller.CreateCategory())
	incomingRoutes.PATCH("/category/:category_id", controller.UpdateCategory())
}
