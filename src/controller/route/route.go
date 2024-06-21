package routes


func InitRoutes(r *gin.RouterGroup, customController controller.CustomControllerInterface) {

	r.GET("/customs/:id", customController.FindById)
	r.POST("/customs/", customController.Create)
}
