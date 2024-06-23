package dao

type CustomDAO interface {
	Create(
		customDomain domain.CustomDomainInterface,
	) (domain.CustomDomainInterface, *rest_err.RestErr)
	FindByID(id string) (domain.CustomDomainInterface, *rest_err.RestErr)
}

type customDAO struct {
	databaseConnection *gorm.DB
}

func NewCustomDAO(database *gorm.DB) CustomDAO {
	return &customDAO{database}
}
