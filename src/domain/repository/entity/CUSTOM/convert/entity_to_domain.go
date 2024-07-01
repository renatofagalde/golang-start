package convert

import (
	"main/src/domain"
	"main/src/domain/repository/entity/custom"
)

func ConvertEntityToDomain(customEntity custom.CustomEntity) domain.CustomDomainInterface {
	domain := domain.NewCustomDomain(
		customEntity.ID,
		customEntity.Custom,
		customEntity.FullName,
		customEntity.Email,
	)

	return domain
}
