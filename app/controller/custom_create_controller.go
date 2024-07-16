package controller 

import (
	"main/app/config/validation"
	"main/app/controller/model/request"
	"main/app/view"
	"net/http"
  "main/app/domain"

	"github.com/gin-gonic/gin"
	toolkit "github.com/renatofagalde/golang-toolkit"
)

func (customController *customControllerInterface) Create(c *gin.Context) {
	var logger toolkit.Logger
	var restErr toolkit.RestErr
	var customRequest request.CustomRequest

	if err := c.ShouldBindJSON(&customRequest); err != nil {
		errRest := validation.ValidateError(err)
		logger.Error("Erro ao validar custom", errRest)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := domain.NewCustomDomain("", customRequest.Custom, customRequest.FullName, customRequest.Email)
  
	result, err := customController.service.Create(domain)
	if err != nil {
		errorMessage := restErr.NewBadRequestError("Erro ao salvar")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(result))
}
