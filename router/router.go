package router

import (
	"github.com/gin-gonic/gin"
)

func StartService() {
	router := gin.Default()
	HomeRoutes(router)
	KonsultasiRoutes(router)
	router.Run()
}
