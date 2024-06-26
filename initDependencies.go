package main

import (
	"main/src/controller/custom"
	"main/src/domain/repository/dao"
	"main/src/domain/service"

	"gorm.io/gorm"
)

func initDependencies(database *gorm.DB) controller.UserControllerInterface {
	r := dao.NewCustomDAO(database)
	domainService := service.NewCustomDomainService(r)
	return controller.NewControllerInterface(domainService)
}
