package main

import (
	"github.com/jayromofo/jayronator/database"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jayromofo/jayronator/routes"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean the kitchen", Completed: false},
	{ID: "2", Item: "Clean the room", Completed: false},
	{ID: "3", Item: "Clean the living room", Completed: false},
}

type Response struct {
	Message string
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()
	db := database.GetClient()
	println(db)
	//router.Use(CORSMiddleware())
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	//router.Use(cors.New(cors.Config{
	//	AllowOrigins: []string{"localhost:3000"}
	//}))
	router.Use(cors.New(config))
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	//router.Use(middleware.Authentication)

	routes.CategoryRoutes(router)
	routes.IngredientRoutes(router)
	routes.MenuRoutes(router)

	router.GET("/")
	router.GET("/api/todos", getTodos)
	router.Run("localhost:3001")
}
