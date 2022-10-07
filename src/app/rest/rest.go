package main

import (
	"context"
	"github.com/dalikewara/ayapingping-go-crud/src/config/constant"
	"github.com/dalikewara/ayapingping-go-crud/src/config/env"
	"github.com/dalikewara/ayapingping-go-crud/src/database/postgresql"
	"github.com/dalikewara/ayapingping-go-crud/src/delivery/rest"
	"github.com/dalikewara/ayapingping-go-crud/src/repository"
	"github.com/dalikewara/ayapingping-go-crud/src/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	ctx := context.Background()

	postgreSQLPool, err := postgresql.ConnectPgxPool(postgresql.ConnectPgxPoolParam{
		Ctx:    ctx,
		Host:   env.PostgreSQLHost,
		Port:   env.PostgreSQLPort,
		User:   env.PostgreSQLUser,
		Pass:   env.PostgreSQLPass,
		DBName: env.PostgreSQLDBName,
		Option: env.PostgreSQLOption,
	})
	if err != nil {
		panic(err)
	}

	timezoneOffset, err := strconv.Atoi(env.TimezoneOffset)
	if err != nil {
		panic(err)
	}
	repo := repository.New(repository.NewParam{
		TimezoneName:   env.TimezoneName,
		TimezoneOffset: timezoneOffset,
		PostgreSQLPool: postgreSQLPool,
	})

	svc := service.New(service.NewParam{
		Repo: repo,
	})

	server := rest.New(rest.NewParam{
		Service: svc,
		Config: &rest.Config{
			Env:                env.AppEnv,
			IsProduction:       env.AppEnv == constant.AppEnvProduction,
			DefaultOkCode:      constant.DefaultOkCode,
			DefaultOkMessage:   constant.DefaultOkMessage,
			DefaultOkStatus:    constant.DefaultOkStatus,
			DefaultNotOkStatus: constant.DefaultNotOkStatus,
		},
		GinClient: gin.Default(),
	})

	server.Gin.RegisterRoutes()
	port, err := strconv.Atoi(env.RESTPort)
	if err != nil {
		panic(err)
	}
	if err = server.Gin.Serve(port); err != nil {
		panic(err)
	}
}
