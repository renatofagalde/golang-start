package repository

import (
	"main/app/domain"
	"main/app/domain/repository/entity"
	"main/app/domain/repository/entity/converter"  
	"fmt"
	"go.uber.org/zap"

	toolkit "github.com/renatofagalde/golang-toolkit"
)

func (customRepository *customRepository) FindByID(id string) (domain.CustomDomainInterface, *toolkit.RestErr) {
	var customEntity entity.CustomEntity
  var logger toolkit.Logger
  var restErr toolkit.RestErr

	err := customRepository.databaseConnection.Where("tenent_id =?", id).First(&customEntity).Error
	if err != nil {
		errorMessage := fmt.Sprintf("custom not found with this ID: %s", id)
		logger.Error(
			fmt.Sprintf("repository.FindById ->  %s", errorMessage),
			err,
			zap.String("journey", "findByID"),
		)
		return nil, restErr.NewNotFoundError(errorMessage)
	}

	logger.Info(
		fmt.Sprintf("repository.FindById ->  %+v", customEntity), zap.String("journey", "findByID"))

	return converter.ConvertEntityToDomain(customEntity), nil
}
