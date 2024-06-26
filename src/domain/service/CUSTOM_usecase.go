package custom

import (
	"main/src/domain"
	"main/src/domain/repository/dao/CUSTOM"

	toolkit "github.com/renatofagalde/golang-toolkit"
)

func NewCustomDomainService(customDAO custom.customDAO) CustomDomainService {
	return &customDomainService{customRepository}
}

type customDomainService struct {
	dao custom.customDAO
}

type CustomDomainService interface {
	Create(domain.customDomainInterface) (domain.customDomainInterface, *toolkit.RestErr)
	FindByID(string) (domain.customDomainInterface, *toolkit.RestErr)
}
