package controller

import (
	"main/src/domain/service/custom"

	"github.com/gin-gonic/gin"
)

type CustomControllerInterface interface {
	Create(c *gin.Context)
	FindById(c *gin.Context)
}

type customControllerInterface struct {
	service custom.CustomDomainService
}

func NewCustomControllerInterface(
	serviceInterface custom.CustomDomainService
) CustomControllerInterface {
	return &customControllerInterface{service: serviceInterface}
}
