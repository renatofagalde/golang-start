package convert

import (
	"main/src/domain"
	"main/src/domain/repository/entity/custom"

)

func ConvertDomainToEntity(domain domain.CustomDomainInterface) *custom.CustomEntity {
	return &custom.CustomEntity{
		ID:       domain.GetID(),
		Custom:   domain.GetCustom(),
		FullName: domain.GetFullName(),
		Email:    domain.GetEmail(),
	}
}
