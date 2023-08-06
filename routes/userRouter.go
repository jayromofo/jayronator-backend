package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/jayromofo/jayronator/controllers"
)

func UserRoutes(incomingRouter *gin.Engine) {
	incomingRouter.GET("/users", controller.GetUsers())
	incomingRouter.GET("/users/:user_id", controller.GetUser())
	incomingRouter.POST("/users/signup", controller.SignUp())
	incomingRouter.POST("/users/login", controller.Login())
	incomingRouter.PATCH("/users/:user_id", controller.UpdateUser())
}
