package controller

import (
	"main/app/domain/service"
	"github.com/gin-gonic/gin"
)

type CustomControllerInterface interface {
	Create(c *gin.Context)
	FindById(c *gin.Context)
}

type customControllerInterface struct {
  service service.CustomDomainService
}

func NewController(serviceInterface service.CustomDomainService) CustomControllerInterface {
	return &customControllerInterface{service: serviceInterface}
}
