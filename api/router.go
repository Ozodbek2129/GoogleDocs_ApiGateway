package api

import (
	"api_gateway/api/handler"
	middleware "api_gateway/api/middlerware"
	"api_gateway/config"
	"log/slog"

	"github.com/gin-gonic/gin"

	_ "api_gateway/api/docs"

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
		users.GET("/profile/:id", h.GetUserProfile)
		users.PUT("/updateUser/:id", h.UpdateUser)
		users.GET("/email/:email", h.GetUserByEmail)
	}

	health := router.Group("/health")
	{
		health.POST("/generate", h.GenerateHealthRecommendations)
		health.GET("/getRealtimeHealthMonitoring/:user_id", h.GetRealtimeHealthMonitoring)
		health.GET("/getDailyHealthSummary/:date", h.GetDailyHealthSummary)
		health.GET("/getWeeklyHealthSummary/:start_date/:end_date", h.GetWeeklyHealthSummary)
	}

	
}
