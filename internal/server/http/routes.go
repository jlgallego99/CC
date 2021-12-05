package http

import (
	"github.com/gin-gonic/gin"
)

func startRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/osts/:obra/:ost", NewOST)

	return router
}
