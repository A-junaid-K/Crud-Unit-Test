package main

import (
	"github.com/testing_ap/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.UserRoutes(router)
	router.Run(":5000")
}
