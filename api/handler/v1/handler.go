package handlers

import (
	"github.com/Asliddin3/api-gateway/config"
	"github.com/Asliddin3/api-gateway/pkg/logger"
	"github.com/Asliddin3/api-gateway/services"
)

type handlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	cfg            config.Config
}

//HandlerV1Config api configuration
type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Cfg            config.Config
}

//New connect to services
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
	}
}
