package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	toolkit "github.com/renatofagalde/golang-toolkit"

	postgres "main/app/config/database/postgres/gorm"
	"main/app/controller/route"
)

func init() {
}

func main() {
	var tools toolkit.Tools
	var logger toolkit.Logger
	config, err := tools.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot read the env file")
	}

	database, err := postgres.NewPostgresGORMConnection(context.Background(), config.DBSource)
	if err != nil {
		log.Fatalf("Error ao conectar no no banco, error=%s", err.Error())
		return
	} else {
		fmt.Println("conexao com sucesso")
	}

	fmt.Println(database)
	controller := initDependencies(database)
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())
	route.InitRoutes(&router.RouterGroup, controller)
	if err := router.Run(config.ServerAddress); err != nil {
		log.Fatal(err)
	}
	logger.Info("Iniciando")
}
