package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quzuu-be/controller"
)

func KonsultasiRoutes(route *gin.Engine) {
	homeRouter := route.Group("/api")
	{
		homeRouter.POST("/create-konsultasi", controller.KonsultasiController)
	}
}
