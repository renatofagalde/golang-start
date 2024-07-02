package main

import (
	"main/src/controller"
	"main/src/domain/repository"
	"main/src/domain/service"

	"gorm.io/gorm"
)

func initDependencies(database *gorm.DB) controller.CustomControllerInterface {
  r := repository.NewCustomRepository(database)
	domainService := service.NewCustomDomainService(r)
	return controller.NewController(domainService)
}
