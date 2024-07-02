package view

import (
	"main/src/controller/model/response"
	"main/src/domain"
)

func ConvertDomainToResponse(domain domain.CustomDomainInterface) response.CustomResponse {
	return response.CustomResponse{
		Custom:   domain.GetCustom(),
		FullName: domain.GetFullName(),
		Email:    domain.GetEmail(),
	}
}
