package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/jayromofo/jayronator/controllers"
)

func RecipeRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/recipes", controller.GetRecipes())
	incomingRoutes.GET("/recipes/:irecipe_id", controller.GetRecipe())
	incomingRoutes.POST("/recipes", controller.CreateRecipe())
	incomingRoutes.PATCH("/recipes/:recipe_id", controller.UpdateRecipe())
	incomingRoutes.DELETE("/recipes/:recipe_id", controller.DeleteRecipe())
}
