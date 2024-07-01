package custom

import (
	"main/src/domain"
  "main/src/domain/service"
  toolkit "github.com/renatofagalde/golang-toolkit"
  "go.uber.org/zap"
  "fmt"
)
func(cd *service.customDomainInterface)	Create(customDomain domain.customDomainInterface) (domain.customDomainInterface, *toolkit.RestErr)
  var logger toolkit.Logger

  logger.Info(fmt.Sprintf("Create: %s sucesso", customDomain), zap.String("journey", "Create"))
	return customDomain
}
