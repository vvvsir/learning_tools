// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/adpter"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/aggregate"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/service"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/conf"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/database/mongo"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/database/redis"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/log"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/repository"
)

// Injectors from wire.go:

func NewApp() (*adpter.Server, error) {
	viper, err := conf.NewViper()
	if err != nil {
		return nil, err
	}
	appConfig, err := conf.NewAppConfigCfg(viper)
	if err != nil {
		return nil, err
	}
	options, err := conf.NewLoggerCfg(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.NewLogger(options)
	if err != nil {
		return nil, err
	}
	config, err := conf.NewMongoConfig(viper)
	if err != nil {
		return nil, err
	}
	client := mongo.NewMongo(config)
	redisConfig, err := conf.NewRedisConfig(viper)
	if err != nil {
		return nil, err
	}
	redisClient, err := redis.NewRedis(redisConfig)
	if err != nil {
		return nil, err
	}
	repositoryRepository := repository.NewRepository(client, redisClient)
	factory := aggregate.NewFactory(repositoryRepository)
	authSrv := service.NewService(repositoryRepository, factory)
	server := adpter.NewSrv(appConfig, logger, authSrv)
	return server, nil
}

// wire.go:

//go:generate wire
var providerSet = wire.NewSet(conf.NewViper, conf.NewAppConfigCfg, conf.NewLoggerCfg, conf.NewRedisConfig, conf.NewMongoConfig, log.NewLogger, redis.NewRedis, mongo.NewMongo, repository.NewRepository, aggregate.NewFactory, service.NewService, adpter.NewSrv)
