package router

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/scan")
	router.POST("/connect")
	router.GET("/weight")

	return router
}
