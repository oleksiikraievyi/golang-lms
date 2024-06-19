package server

import (
	"lms/handlers"

	docs "lms/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func NewRouter(cfg *Config) *gin.Engine {
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Title = "Lead Management System"
	docs.SwaggerInfo.Description = "This is a simple API for managing clients and leads"

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	clientHandler := handlers.NewClientHandler()

	router.GET("/clients", clientHandler.GetClients)
	router.GET("/clients/:id", clientHandler.GetClient)
	router.POST("/clients", clientHandler.CreateClient)
	router.PUT("/lead", clientHandler.AssignLeadToClient)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
