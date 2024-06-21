package view

func ConvertDomainToResponse(domain domain.CustomDomainInterface) model.CustomResponse {
	return response.CustomResponse{
		Custom:                   domain.GetCustom(),
		FullName:             domain.GetFullName(),
		Email:       domain.GetEmail()
	}
}