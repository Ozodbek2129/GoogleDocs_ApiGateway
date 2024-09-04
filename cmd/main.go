package main

import (
	"api_gateway/api"
	"api_gateway/api/handler"
	"api_gateway/casbin"
	"api_gateway/config"
	"api_gateway/pkg/logger"
	"api_gateway/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("API Gateway started successfully!")
	logger := logger.NewLogger()
	logger.Info("API Gateway started successfully!")

	enforcer, err := casbin.CasbinEnforcer(logger)
	if err != nil {
        log.Println("Error initializing casbin enforcer", "error", err.Error())
		logger.Error("Error initializing enforcer", "error", err.Error())
		return
    }

	config := config.Load()
	serviceManager, err := service.NewServiceManager()
	if err != nil {
		log.Println("Error initializing service manager", "error", err.Error())
		logger.Error("Error initializing service manager", "error", err.Error())
		return
	}


	handler := handler.NewHandler(serviceManager.UserService(), serviceManager.Docsservice(), logger, enforcer)
	controller := api.NewController(gin.Default())
	controller.SetupRoutes(*handler, logger)
	controller.StartServer(config)

}