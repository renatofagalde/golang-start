package controller


type CustomControllerInterface interface {
	Create(c *gin.Context)
	FindById(c *gin.Context)
}

type customControllerInterface struct {
	service service.CustomDomainService
}

func NewCustomControllerInterface(
	serviceInterface service.CustomDomainService,
) CustomControllerInterface {
	return &customControllerInterface{service: serviceInterface}
}
