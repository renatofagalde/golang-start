package controller 

import (
	"main/src/config/validation"
	"main/src/controller/model/request"
	"main/src/view/convert"
  "main/src/controller"
	"net/http"

	"github.com/gin-gonic/gin"
	toolkit "github.com/renatofagalde/golang-toolkit"
)

func (customController *controller.customControllerInterface) Create(c *gin.Context) {
	var logger toolkit.Logger
	var restErr toolkit.RestErr
	var customRequest request.CustomRequest

	if err := c.ShouldBindJSON(&customRequest); err != nil {
		errRest := validation.ValidatecustomError(err)
		logger.Error("Erro ao validar custom", errRest)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewCustomDomain(nil, customRequest.Custom, customRequest.fullName, customRequest.email)

	result, err := customController.service.Create(domain)
	if err != nil {
		errorMessage := restErr.NewBadRequestError("Erro ao salvar")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	c.JSON(http.StatusOK, convert.ConvertDomainToResponse(result))
}
