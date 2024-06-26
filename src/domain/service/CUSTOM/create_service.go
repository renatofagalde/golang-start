package custom

import (
	"main/src/domain"
  toolkit "github.com/renatofagalde/golang-toolkit"
)
func(cd *customDomainInterface)	CreateService(customDomain domain.customDomainInterface) (domain.customDomainInterface, *toolkit.RestErr)
  var logger toolkit.Logger

	logger.Info("init create service")
	return customDomain
}
