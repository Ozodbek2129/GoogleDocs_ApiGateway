package api

import (
	"api_gateway/api/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "api_gateway/api/docs"
)

// @title        E-Commerce API
// @version      1.0
// @description  This is an API for e-commerce platform.
// @termsOfService http://swagger.io/terms/
// @contact.name  API Support
// @contact.email support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host          localhost:9876
// @BasePath      /
func NewRouter(h *handler.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user := router.Group("/user")
	{
		user.GET("getbyuser/:email",h.GetUSerByEmail)
		user.PUT("/update_user",h.UpdateUser)
		user.DELETE("/delete_user/:id",h.DeleteUser)
	}
	return router
}
