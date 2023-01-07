package routers

import (
	"a_simple_api_go/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", controllers.Test)
	r.GET("/address", controllers.Postal)

	return r
}
