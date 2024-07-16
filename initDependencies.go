package main

import (
	"main/app/controller"
	"main/app/domain/repository"
	"main/app/domain/service"

	"gorm.io/gorm"
)

func initDependencies(database *gorm.DB) controller.CustomControllerInterface {
  r := repository.NewCustomRepository(database)
	domainService := service.NewCustomDomainService(r)
	return controller.NewController(domainService)
}
