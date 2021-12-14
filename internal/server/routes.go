package server

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Host string
	Port string
}

func SetupRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	gin.DefaultWriter = os.Stdout

	router.POST("/osts/:obra", newOST)

	router.GET("/osts", OSTs)
	router.GET("/osts/:obra/:ostid", getOST)

	router.NoRoute(NoRoute)

	return router
}
