package dao

import (
	"fmt"
	"main/src/domain"
	"main/src/domain/repository/entity/custom/convert"
  "main/src/domain/repository/entity/custom"


	toolkit "github.com/renatofagalde/golang-toolkit"
	"go.uber.org/zap"
)

func (customRepository *customDAO) Create(customDomain domain.CustomDomainInterface) (domain.CustomDomainInterface, *toolkit.RestErr) {
	var logger toolkit.Logger
	logger.Info("init create custom repository", zap.String("journey", "create"))

	toSave := &custom.CustomEntity{
		ID:       "1",
		Custom:   customDomain.Custom(),
		FullName: customDomain.FullName(),
		Email:    customDomain.GetEmail(),
	}

	customRepository.databaseConnection.Create(&toSave)

	logger.Info(fmt.Sprintf("custom entity object %+v", toSave))
	return converter.ConvertEntityToDomain(*toSave), nil
}
