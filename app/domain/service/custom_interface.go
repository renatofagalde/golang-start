package service

import (
	"main/app/domain"
	"main/app/domain/repository"

	toolkit "github.com/renatofagalde/golang-toolkit"
)

func NewCustomDomainService(customRepository repository.CustomRepository) CustomDomainService{
	return &customDomainService{customRepository}
}

type customDomainService struct {
	customRepository repository.CustomRepository
}

type CustomDomainService interface {
	Create(domain.CustomDomainInterface) (domain.CustomDomainInterface, *toolkit.RestErr)
	FindByID(string) (domain.CustomDomainInterface, *toolkit.RestErr)
}
