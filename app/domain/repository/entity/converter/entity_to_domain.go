package converter

import (
	"main/app/domain"
	"main/app/domain/repository/entity"
)

func ConvertEntityToDomain(customEntity entity.CustomEntity) domain.CustomDomainInterface {
	domain := domain.NewCustomDomain(
		customEntity.ID,
		customEntity.Custom,
		customEntity.FullName,
		customEntity.Email,
	)

	return domain
}
