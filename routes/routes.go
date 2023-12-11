package routes

import (
	"github.com/testing_ap/controllers"
	"github.com/testing_ap/database"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	db := database.InitDatabase()
	router.POST("/", controllers.CreateUser(db))
	router.GET("/:id", controllers.GetUserById(db))
	router.PUT("/:id",controllers.UpdateUser(db))
	router.DELETE("/:id",controllers.DeleteUser(db))
}
