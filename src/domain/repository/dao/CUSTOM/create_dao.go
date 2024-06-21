package repository

import (
	"backend-custom/src/config/logger"
	"backend-custom/src/config/rest_err"
	"backend-custom/src/model"
	"backend-custom/src/model/repository/entity"
	"backend-custom/src/model/repository/entity/converter"
	"fmt"
	"go.uber.org/zap"
)

func (sr *customRepository) Create(customDomain model.customDomainInterface) (model.customDomainInterface, *rest_err.RestErr) {
	logger.Info("init create custom repository", zap.String("journey", "create"))

	toSave := &entity.customEntity{Title: customDomain.GetTitle()}

	sr.databaseConnection.Create(&toSave)

	logger.Info(fmt.Sprintf("custom entity object %+v", toSave))
	return converter.ConvertEntityToDomain(*toSave), nil
}
