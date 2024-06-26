package custom

import (
	"main/src/domain"
	"main/src/domain/repository/dao"
	"main/src/domain/repository/entity"

	toolkit "github.com/renatofagalde/golang-toolkit"
)

func (customRepository *dao.customDAO) FindByID(
	id string,
) (domain.customDomainInterface, *toolkit.RestErr) {
	var customEntity entity.customEntity

	err := customRepository.databaseConnection.Where("tenent_id =?", id).First(&customEntity).Error
	if err != nil {
		errorMessage := fmt.Sprintf("custom not found with this ID: %s", id)
		logger.Error(
			fmt.Sprintf("repository.FindById ->  %s", errorMessage),
			err,
			zap.String("journey", "findByID"),
		)
		return nil, rest_err.NewNotFoundError(errorMessage)
	}

	logger.Info(
		fmt.Sprintf("repository.FindById ->  %+v", customEntity), zap.String("journey", "findByID"))

	return converter.ConvertEntityToDomain(customEntity), nil
}
