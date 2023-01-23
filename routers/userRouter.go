package routes 

import (
	"github.com/gin-gonic/gin"
	controller "resturant-mg/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingROutes.GET("/users", controller.GetUsers())
	incomingROutes.GET("/users/:user_id", controller.Get)
	incomingROutes.POST("/users/signup", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())
}