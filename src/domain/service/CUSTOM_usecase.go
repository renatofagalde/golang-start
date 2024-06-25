package service

func NewCustomDomainService(customDAO dao.customDAO) customDomainService {
	return &customDomainService{customRepository}
}

type customDomainService struct {
	customDAO dao.customDAO
}
type customDomainService interface {
	CreateService(domain.customDomainInterface) (domain.customDomainInterface, *rest_err.RestErr)
	FindcustomByIDService(id string) (domain.customDomainInterface, *rest_err.RestErr)
}
