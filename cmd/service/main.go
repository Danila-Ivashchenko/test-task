package main

import (

	// "go-kafka/internal/adapters/kafka"

	"go-kafka/internal/adapters/hasher"
	"go-kafka/internal/adapters/http-server/handler"
	"go-kafka/internal/adapters/http-server/middleware"
	"go-kafka/internal/adapters/http-server/server"
	"go-kafka/internal/adapters/kafka"

	// "go-kafka/internal/adapters/kafka"
	user_stor "go-kafka/internal/adapters/storage/user"
	"go-kafka/internal/domain/service"
	"go-kafka/internal/pkg/enricher"
	"go-kafka/internal/pkg/validator"
	"go-kafka/pkg/config"
	"go-kafka/pkg/logger"
	"go-kafka/pkg/psql"
	"go-kafka/pkg/redis"
)

func main() {
	err := config.LoadEnv(".env")
	if err != nil {
		panic(err)
	}
	cfg := config.GetConfig()
	logger := logger.SetupLogger(cfg.GetEnv())
	logger.Debug("debug is avalible")

	e := enricher.New()
	psqlClient := psql.NewPostgresClient(cfg)
	s := user_stor.NewUserStorage(psqlClient)
	redisClient := redis.NewRedisClient(cfg)
	hasher := hasher.New(cfg, redisClient)
	userService := service.NewUserService(s, e, validator.New(), hasher, logger)

	middleware := middleware.New(logger)
	handler := handler.New(cfg, userService)
	server := server.New(cfg, handler, middleware)
	go func() {
		server.Run()
	}()

	kafka, err := kafka.New(cfg, userService, logger)
	if err != nil {
		panic(err)
	}
	kafka.Consume()
}
