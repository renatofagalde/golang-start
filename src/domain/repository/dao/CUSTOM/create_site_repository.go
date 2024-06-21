package repository

import (
	"backend-site/src/config/logger"
	"backend-site/src/config/rest_err"
	"backend-site/src/model"
	"backend-site/src/model/repository/entity"
	"backend-site/src/model/repository/entity/converter"
	"fmt"
	"go.uber.org/zap"
)

func (sr *siteRepository) Create(siteDomain model.SiteDomainInterface) (model.SiteDomainInterface, *rest_err.RestErr) {
	logger.Info("init create site repository", zap.String("journey", "create"))

	toSave := &entity.SiteEntity{Title: siteDomain.GetTitle()}

	sr.databaseConnection.Create(&toSave)

	logger.Info(fmt.Sprintf("site entity object %+v", toSave))
	return converter.ConvertEntityToDomain(*toSave), nil
}
