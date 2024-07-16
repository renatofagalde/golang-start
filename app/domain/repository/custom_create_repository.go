package repository

import (
	"fmt"
	"main/app/domain"
  "main/app/domain/repository/entity"
  "main/app/domain/repository/entity/converter"
	toolkit "github.com/renatofagalde/golang-toolkit"
	"go.uber.org/zap"
)

func (customRepository *customRepository) Create(customDomain domain.CustomDomainInterface) (domain.CustomDomainInterface, *toolkit.RestErr) {
	var logger toolkit.Logger
	logger.Info("init create custom repository", zap.String("journey", "create"))

	toSave := &entity.CustomEntity{
		Custom:   customDomain.GetCustom(),
		FullName: customDomain.GetFullName(),
		Email:    customDomain.GetEmail(),
	}

	customRepository.databaseConnection.Create(&toSave)

	logger.Info(fmt.Sprintf("custom entity object %+v", toSave))
	return converter.ConvertEntityToDomain(*toSave), nil
}
