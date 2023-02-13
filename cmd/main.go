package main

import (
	"github.com/Asliddin3/api-gateway/api"
	"github.com/Asliddin3/api-gateway/config"
	"github.com/Asliddin3/api-gateway/pkg/logger"
	"github.com/Asliddin3/api-gateway/services"
	r "github.com/Asliddin3/api-gateway/storage/redis"
	"github.com/gomodule/redigo/redis"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api-gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}
	pool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "redisdb:6379")
		},
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
		Redis:          r.NewRedisRepo(pool),
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
