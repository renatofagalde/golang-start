package controller
import (
	"encoding/json"
	"fmt"
	"main/src/view/convert"
	"net/http"
  "main/src/controller"

	"github.com/gin-gonic/gin"
	toolkit "github.com/renatofagalde/golang-toolkit"
	"go.uber.org/zap"
)

func (customController *controller.customControllerInterface) FindById(c *gin.Context) {
	var logger toolkit.Logger

	logger.Info("controller.FindById", zap.String("journey", "findByID"))
	id := c.Param("id")
	out, _ := json.Marshal(c)
	logger.Info(fmt.Sprintf("context %+v", out), zap.String("journey", "findByID"))
	origin := c.Request.Header.Get("Origin")
	logger.Info(fmt.Sprintf("origin: %s", origin))

	customDomain, err := customController.service.FindByID(id)
	if err != nil {
		c.JSON(err.Code, err)
		message := "Error retriving by ID"
		logger.Error(message, err)
		return
	}
	result := convert.ConvertDomainToResponse(customDomain)
	logger.Info(
		fmt.Sprintf("init FindUserByID find_controller successfuly %+v", result),
		zap.String("journey", "findByID"),
	)
	c.JSON(http.StatusOK, result)
}
