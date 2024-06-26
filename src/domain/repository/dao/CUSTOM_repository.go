package dao

import (
	"main/src/domain"

	toolkit "github.com/renatofagalde/golang-toolkit"
	"gorm.io/gorm"
)

type CustomDAO interface {
	Create(
		customDomain domain.CustomDomainInterface,
	) (domain.CustomDomainInterface, *toolkit.RestErr)
	FindByID(id string) (domain.CustomDomainInterface, *toolkit.RestErr)
}

type customDAO struct {
	databaseConnection *gorm.DB
}

func NewCustomDAO(database *gorm.DB) CustomDAO {
	return &customDAO{database}
}
