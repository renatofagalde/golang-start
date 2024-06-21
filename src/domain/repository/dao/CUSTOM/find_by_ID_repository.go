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

func (siteRepository *siteRepository) FindByID(id string) (model.SiteDomainInterface, *rest_err.RestErr) {

	var siteEntity entity.SiteEntity

	err := siteRepository.databaseConnection.Where("tenent_id =?", id).First(&siteEntity).Error
	if err != nil {
		errorMessage := fmt.Sprintf("Site not found with this ID: %s", id)
		logger.Error(fmt.Sprintf("repository.FindById ->  %s", errorMessage), err, zap.String("journey", "findByID"))
		return nil, rest_err.NewNotFoundError(errorMessage)
	}

	logger.Info(
		fmt.Sprintf("repository.FindById ->  %+v", siteEntity), zap.String("journey", "findByID"))

	return converter.ConvertEntityToDomain(siteEntity), nil
}
