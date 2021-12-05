package http

import (
	"github.com/gin-gonic/gin"
)

func startRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/osts/:obra/:ost", NewOST)
	router.GET("/osts/:obra/:ost", GetOST)

	return router
}
