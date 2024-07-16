package service 

import (
	"main/app/domain"
  toolkit "github.com/renatofagalde/golang-toolkit"
  "go.uber.org/zap"
  "fmt"
)
func(cd *customDomainService)	Create(customDomain domain.CustomDomainInterface) (domain.CustomDomainInterface, *toolkit.RestErr){
  var logger toolkit.Logger

  logger.Info(fmt.Sprintf("Create: %s sucesso", customDomain), zap.String("journey", "Create"))
	return customDomain, nil
}
