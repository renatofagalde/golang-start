package repository

import (
	"main/src/domain"

	toolkit "github.com/renatofagalde/golang-toolkit"
	"gorm.io/gorm"
)

type CustomRepository interface {
	Create(customDomain domain.CustomDomainInterface) (domain.CustomDomainInterface, *toolkit.RestErr)
	FindByID(id string) (domain.CustomDomainInterface, *toolkit.RestErr)
}

type customRepository struct {
	databaseConnection *gorm.DB
}

func NewCustomRepository(database *gorm.DB) CustomRepository {
	return &customRepository{database}
}
