package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	public := router.Group("/api/v1")

	public.GET("/patients", controllers.GetAllPatients)

	return router
}
