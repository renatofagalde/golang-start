package controller

import (
	"net/http"
)

func (customController *customControllerInterface) Create(c *gin.Context) {

	var customRequest model.CustomRequest

	if err := c.ShouldBindJSON(&siteRequest); err != nil {
		errRest := validation.ValidateError(err)
		logger.Error("Erro ao validar site", errRest)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewSiteDomain(siteRequest.Title, "", "", "", "")

	result, err := siteController.service.Create(domain)
	if err != nil {
		c.JSON(err.Code, err)
		logger.Error("Erro ao chamar o create ", err)
		return
	}
	logger.Info("init create userController")

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(result))
}
