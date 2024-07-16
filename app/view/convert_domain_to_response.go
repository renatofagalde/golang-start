package view

import (
	"main/app/controller/model/response"
	"main/app/domain"
)

func ConvertDomainToResponse(domain domain.CustomDomainInterface) response.CustomResponse {
	return response.CustomResponse{
		Custom:   domain.GetCustom(),
		FullName: domain.GetFullName(),
		Email:    domain.GetEmail(),
	}
}
