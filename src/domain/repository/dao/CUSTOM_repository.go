package dao


type CustomDAO interface {
	Create(customDomain domain.CusomDomainInterface) (domain.CusomDomainInterface, *rest_err.RestErr)
	FindByID(id string) (domain.CusomDomainInterface, *rest_err.RestErr)
}

type customDAO struct {
	databaseConnection *gorm.DB
}

func NewCustomDAO(database *gorm.DB) CustomDAO {
	return &customDAO{database}
}



