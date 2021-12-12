package server

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/osts/:obra", newOST)

	router.GET("/osts", OSTs)
	router.GET("/osts/:obra/:ostid", getOST)

	router.NoRoute(NoRoute)

	return router
}
