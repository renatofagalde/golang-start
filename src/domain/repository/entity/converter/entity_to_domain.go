package converter

import (
	"main/src/domain"
	"main/src/domain/repository/entity"
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
