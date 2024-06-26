package convert

import (
	"main/src/domain"
	"main/src/domain/repository/entity/custom"
)

func ConvertDomainToEntity(domain domain.CustomDomainInterface) *entity.CustomEntity {
	return &entity.CustomEntity{
		ID:       domain.GetID(),
		Custom:   domain.GetCustom(),
		FullName: domain.GetFullName(),
		Email:    domain.GetEmail(),
	}
}
