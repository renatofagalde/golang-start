package service

import (
	"main/src/domain"
	"main/src/domain/repository/dao/CUSTOM"

	toolkit "github.com/renatofagalde/golang-toolkit"
)

func NewCustomDomainService(customDAO custom.customDAO) customDomainService {
	return &customDomainService{customRepository}
}

type customDomainService struct {
	customDAO custom.customDAO
}
type customDomainService interface {
	CreateService(domain.customDomainInterface) (domain.customDomainInterface, *toolkit.RestErr)
	FindByIDService(id string) (domain.customDomainInterface, *toolkit.RestErr)
}
