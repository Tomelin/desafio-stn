package main

import (
	"context"
	"log"

	"github.com/Tomelin/desafio-stn/src/configs"
	_ "github.com/Tomelin/desafio-stn/src/docs/swagger"
	"github.com/Tomelin/desafio-stn/src/internal/core/repository"
	"github.com/Tomelin/desafio-stn/src/internal/core/service"
	"github.com/Tomelin/desafio-stn/src/internal/infra/handler/kube"
	"github.com/Tomelin/desafio-stn/src/internal/infra/handler/webserver"
	ra "github.com/Tomelin/desafio-stn/src/pkg/rest-api"
	"github.com/Tomelin/desafio-stn/src/pkg/utils"
)

// @title 				kubernetes-manager
// @version				1.0
// @description		Microservice to kubernetes manage

// // @securityDefinitions.basic BasicAuth
// // @in header
// // @name Authorization

// @contact.email rafael.tomelin@gmail.com

// schemes		http https
// @BasePath	/api
func main() {

	// load config
	configServer, err := configs.LoadConfigs()
	if err != nil {
		log.Fatal(err)
	}

	// start kubernetes connection
	kubeConnection, err := kube.NewKubeConnection(configServer.Kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// load http server
	router := ra.NewRestAPI(configServer.Webserver.Http2Enabled, configServer.Webserver.Mode)

	// start the namespace
	rnamespace := repository.NewRepositoryNamespace(kubeConnection)
	snamespace := service.NewServiceNamespace(rnamespace)
	webserver.NewNamespaceWebServer(router.Group, snamespace)

	// start the deployment
	rdeployment := repository.NewDeploymentRepository(kubeConnection)
	sdeployment := service.NewDeploymentService(rdeployment)
	webserver.NewDeploymentWebServer(router.Group, sdeployment)

	t := utils.Timeout{
		Millisecond: configServer.Timeout,
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, "timeout", t)
	_, err = snamespace.Count(ctx)
	_, err = snamespace.GetAll(ctx)

	// start webserver
	router.Run(router.Route.Handler(), configServer.Webserver.Listen, configServer.Webserver.Port)

	log.Fatal()
}
