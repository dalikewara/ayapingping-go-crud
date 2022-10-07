package postgresql

import (
	"context"
	"errors"
	"fmt"
	"github.com/dalikewara/pgxpoolgo"
)

type ConnectPgxPoolParam struct {
	Ctx    context.Context
	Host   string
	Port   string
	User   string
	Pass   string
	DBName string
	Option string
}

// ConnectPgxPool connects to PostgreSQL database connection pool using pgxpool.
func ConnectPgxPool(param ConnectPgxPoolParam) (pgxpoolgo.Pool, error) {
	connString := "postgres://"

	if param.User != "" && param.Pass != "" {
		connString += param.User + ":" + param.Pass
	} else if param.User != "" && param.Pass == "" {
		connString += param.User
	}
	if param.Host != "" && param.Port != "" {
		connString += "@" + param.Host + ":" + param.Port
	} else if param.Host != "" && param.Port == "" {
		connString += "@" + param.Host + ":5432"
	}
	if param.DBName != "" {
		connString += "/" + param.DBName
	}
	if param.Option != "" {
		connString += "?" + param.Option
	}

	connConfig, err := pgxpoolgo.ParseConfig(connString)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(ErrParseConfig, err.Error()))
	}
	pool, err := pgxpoolgo.ConnectConfig(param.Ctx, connConfig)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(ErrConnect, err.Error()))
	}
	err = pool.Ping(param.Ctx)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(ErrPing, err.Error()))
	}
	return pool, nil
}
