package custom

func (customRepository *customRepository) FindByID(id string) (domain.customDomainInterface, *rest_err.RestErr) {

	var customEntity entity.customEntity

	err := customRepository.databaseConnection.Where("tenent_id =?", id).First(&customEntity).Error
	if err != nil {
		errorMessage := fmt.Sprintf("custom not found with this ID: %s", id)
		logger.Error(fmt.Sprintf("repository.FindById ->  %s", errorMessage), err, zap.String("journey", "findByID"))
		return nil, rest_err.NewNotFoundError(errorMessage)
	}

	logger.Info(
		fmt.Sprintf("repository.FindById ->  %+v", customEntity), zap.String("journey", "findByID"))

	return converter.ConvertEntityToDomain(customEntity), nil
}
