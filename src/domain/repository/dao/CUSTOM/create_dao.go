package custom

import (
	"fmt"
	"main/src/domain"
	"main/src/domain/repository/dao"
	"main/src/domain/repository/entity"

	toolkit "github.com/renatofagalde/golang-toolkit"
	"go.uber.org/zap"
)

func (customRepository *dao.customDAO) Create(
	customDomain domain.CustomDomainInterface,
) (domain.CustomDomainInterface, *toolkit.RestErr) {
	logger.Info("init create custom repository", zap.String("journey", "create"))

	toSave := &entity.customEntity{
		ID:       1,
		Custom:   customDomain.Custom(),
		FullName: customDomain.FullName(),
		Email:    customDomain.GetEmail(),
	}

	customRepository.databaseConnection.Create(&toSave)

	logger.Info(fmt.Sprintf("custom entity object %+v", toSave))
	return converter.ConvertEntityToDomain(*toSave), nil
}
