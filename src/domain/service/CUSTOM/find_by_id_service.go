package custom

import (
	"main/src/domain"
  "main/src/domain/service"

	toolkit "github.com/renatofagalde/golang-toolkit"
)

func (cd *service.customDomainInterface) FindByID(
	id string,
) (domain.customDomainInterface, *toolkit.RestErr) {
	logger.Info("init FindcustomByIDService ")
	return cd.customDAO.FindCustomByID(id)
}
