package controller

import (
	"backend-site/src/config/logger"
	"backend-site/src/config/validation"
	"backend-site/src/controller/model/request"
	"backend-site/src/model"
	"backend-site/src/view"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (siteController *siteControllerInterface) Create(c *gin.Context) {

	var siteRequest request.SiteRequest

	if err := c.ShouldBindJSON(&siteRequest); err != nil {
		errRest := validation.ValidateSiteError(err)
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
