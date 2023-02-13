package api

import (
	_ "github.com/Asliddin3/api-gateway/api/docs" //swag
	v1 "github.com/Asliddin3/api-gateway/api/handler/v1"
	"github.com/Asliddin3/api-gateway/config"
	"github.com/Asliddin3/api-gateway/pkg/logger"
	"github.com/Asliddin3/api-gateway/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// New ...
// @title           Rest api
// @version         1.0
// @description     This is tz api
// @termsOfService  not much usefull

// @contact.name   Asliddin
// @contact.url    https://t.me/asliddindeh
// @contact.email  asliddinvstalim@gmail.com

// @host      localhost:8070
// @BasePath  /v1

// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowCredentials: true,
		AllowOrigins:     []string{},
	}))

	api := router.Group("/v1")
	// User routers
	api.POST("/user", handlerV1.Create)
	api.GET("/user/list", handlerV1.FindAll)
	api.GET("/user/:id", handlerV1.FindByID)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
