package convert

import (
	"main/src/controller/custom/model"
	"main/src/domain"
)

func ConvertDomainToResponse(domain domain.CustomDomainInterface) model.CustomResponse {
	return model.CustomResponse{
		Custom:   domain.GetCustom(),
		FullName: domain.GetFullName(),
		Email:    domain.GetEmail(),
	}
}
