package custom

import (
	"main/src/domain"

	toolkit "github.com/renatofagalde/golang-toolkit"
)

func (cd *customDomainInterface) FindByID(
	id string,
) (domain.customDomainInterface, *toolkit.RestErr) {
	logger.Info("init FindcustomByIDService ")
	return cd.customDAO.FindCustomByID(id)
}
