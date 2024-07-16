package route

import (
	"main/app/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, customController controller.CustomControllerInterface) {
	r.GET("/customs/:id", customController.FindById)
	r.POST("/customs/", customController.Create)
}
