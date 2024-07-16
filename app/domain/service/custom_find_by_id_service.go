package service

import (
	"main/app/domain"

	toolkit "github.com/renatofagalde/golang-toolkit"
)

func (cd *customDomainService) FindByID(id string) (domain.CustomDomainInterface, *toolkit.RestErr) {
	var logger toolkit.Logger

  logger.Info("init FindcustomByIDService ")
	return cd.customRepository.FindByID(id)
}
