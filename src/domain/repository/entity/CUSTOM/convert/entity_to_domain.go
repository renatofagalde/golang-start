

func ConvertEntityToDomain(customEntity entity.CustomEntity) model.CustomDomainInterface{
domain := domain.NewCustomDomain(customEntity.ID,customEntity.Custom,customEntity.FullName,customEntity.Email)
	logger.Info(
		fmt.Sprintf("convertEntityToDomain ->  %+v", domain), zap.String("journey", "findByID"))

	return domain
}
