package custom

import (
	"main/domain/service"

	toolkit "github.com/renatofagalde/golang-toolkit"
)

func (cd *customDomainService) FindByIDService(
	id string,
) (domain.customDomainInterface, *toolkit.RestErr) {
	logger.Info("init FindcustomByIDService ")
	return cd.customDAO.FindCustomByID(id)
}
