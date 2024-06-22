package custom

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (customController *customControllerInterface) Create(c *gin.Context) {
	var customRequest request.customRequest

	if err := c.ShouldBindJSON(&customRequest); err != nil {
		errRest := validation.ValidatecustomError(err)
		logger.Error("Erro ao validar custom", errRest)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewcustomDomain(customRequest.Title, "", "", "", "")

	result, err := customController.service.Create(domain)
	if err != nil {
		c.JSON(err.Code, err)
		logger.Error("Erro ao chamar o create ", err)
		return
	}
	logger.Info("init create userController")

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(result))
}
