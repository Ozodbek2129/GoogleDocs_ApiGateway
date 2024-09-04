package api

import (
	"api_gateway/api/handler"
	"api_gateway/api/middleware"
	"api_gateway/config"
	"log/slog"

	"github.com/gin-gonic/gin"

	// _ "api_gateway/api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Controller interface {
	SetupRoutes(handler.Handler, *slog.Logger)
	StartServer(config.Config) error
}

type controllerImpl struct {
	Port   string
	Router *gin.Engine
}

func NewController(router *gin.Engine) Controller {
	return &controllerImpl{Router: router}
}

func (c *controllerImpl) StartServer(cfg config.Config) error {
	c.Port = cfg.API_GATEWAY
	return c.Router.Run(c.Port)
}

// @title Api Gateway
// @version 1.0
// @description This is a sample server for Api-gateway Service
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api/v1
// @schemes http
func (c *controllerImpl) SetupRoutes(h handler.Handler, logger *slog.Logger) {
	c.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router := c.Router.Group("/api")
	router.Use(middleware.Check)
	router.Use(middleware.CheckPermissionMiddleware(h.Enforcer))

	users := router.Group("/user")
	{
		users.GET("/getbyuser/:email", h.GetUSerByEmail)
		users.PUT("/update_user", h.UpdateUser)
		users.DELETE("/delete_user/:id", h.DeleteUser)
	}

	docs := router.Group("/docs")
	{
		docs.POST("/createDocument", h.CreateDocument)
		docs.GET("/SearchDocument", h.SearchDocument)
		docs.GET("/GetAllDocuments", h.GetAllDocuments)
		docs.PUT("/UpdateDocument", h.UpdateDocument)
		docs.DELETE("/DeleteDocument", h.DeleteDocument)
		docs.POST("/ShareDocument", h.ShareDocument)
	}

	version := router.Group("/version")
	{
		version.GET("/GetAllVersions", h.GetAllVersions)
		version.PUT("/RestoreVersion", h.RestoreVersion)
	}
}
