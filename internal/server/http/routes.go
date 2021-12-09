package http

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/osts/:obra", newOST)
	router.GET("/osts/:obra/:ostid", getOST)

	return router
}
