//go:build integration

package postgresql_test

import (
	"context"
	"fmt"
	"github.com/dalikewara/ayapingping-go-crud/src/database/postgresql"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"testing"
)

var validPgxPoolConnection = postgresql.ConnectPgxPoolParam{
	Host:   "localhost",
	Port:   "5432",
	User:   "root",
	Pass:   "p455w0rd",
	DBName: "test",
}

func TestConnectPgxPool_OK(t *testing.T) {
	t.Run("normal case", func(t *testing.T) {
		ctx := context.Background()

		postgreSQLPool, err := postgresql.ConnectPgxPool(postgresql.ConnectPgxPoolParam{
			Ctx:    ctx,
			Host:   validPgxPoolConnection.Host,
			Port:   validPgxPoolConnection.Port,
			User:   validPgxPoolConnection.User,
			Pass:   validPgxPoolConnection.Pass,
			DBName: validPgxPoolConnection.DBName,
		})

		assert.Nil(t, err)
		assert.ObjectsAreEqual(&pgxpool.Pool{}, postgreSQLPool)
	})

	t.Run("no port, default 5432 used", func(t *testing.T) {
		ctx := context.Background()

		postgreSQLPool, err := postgresql.ConnectPgxPool(postgresql.ConnectPgxPoolParam{
			Ctx:    ctx,
			Host:   validPgxPoolConnection.Host,
			User:   validPgxPoolConnection.User,
			Pass:   validPgxPoolConnection.Pass,
			DBName: validPgxPoolConnection.DBName,
		})

		assert.Nil(t, err)
		assert.ObjectsAreEqual(&pgxpool.Pool{}, postgreSQLPool)
	})

	t.Run("with option", func(t *testing.T) {
		ctx := context.Background()

		postgreSQLPool, err := postgresql.ConnectPgxPool(postgresql.ConnectPgxPoolParam{
			Ctx:    ctx,
			Host:   validPgxPoolConnection.Host,
			User:   validPgxPoolConnection.User,
			Pass:   validPgxPoolConnection.Pass,
			DBName: validPgxPoolConnection.DBName,
			Option: "pool_max_conns=10",
		})

		assert.Nil(t, err)
		assert.ObjectsAreEqual(&pgxpool.Pool{}, postgreSQLPool)
	})
}

func TestConnectPgxPool_ErrParseConfig(t *testing.T) {
	t.Run("no host & port", func(t *testing.T) {
		ctx := context.Background()

		postgreSQLPool, err := postgresql.ConnectPgxPool(postgresql.ConnectPgxPoolParam{
			Ctx:    ctx,
			User:   "root",
			Pass:   "p455w0rd",
			DBName: "test",
		})

		assert.Nil(t, postgreSQLPool)
		assert.NotNil(t, err)
		assert.EqualError(t, err, fmt.Sprintf(postgresql.ErrParseConfig, "cannot parse `postgres://root:p455w0rd/test`: failed to parse as URL (parse \"postgres://root:p455w0rd/test\": invalid port \":p455w0rd\" after host)"))
	})
}

func TestConnectPgxPool_ErrConnect(t *testing.T) {
	t.Run("authentication failed", func(t *testing.T) {
		ctx := context.Background()

		postgreSQLPool, err := postgresql.ConnectPgxPool(postgresql.ConnectPgxPoolParam{
			Ctx:    ctx,
			Host:   "localhost",
			Port:   "5432",
			User:   "root",
			DBName: "test",
		})

		assert.Nil(t, postgreSQLPool)
		assert.NotNil(t, err)
		assert.EqualError(t, err, fmt.Sprintf(postgresql.ErrConnect, "failed to connect to `host=localhost user=root database=test`: failed SASL auth (FATAL: password authentication failed for user \"root\" (SQLSTATE 28P01))"))
	})

	t.Run("database doesn't exists", func(t *testing.T) {
		ctx := context.Background()

		postgreSQLPool, err := postgresql.ConnectPgxPool(postgresql.ConnectPgxPoolParam{
			Ctx:    ctx,
			Host:   "localhost",
			Port:   "5432",
			User:   "root",
			DBName: "testT",
		})

		assert.Nil(t, postgreSQLPool)
		assert.NotNil(t, err)
		assert.EqualError(t, err, fmt.Sprintf(postgresql.ErrConnect, "failed to connect to `host=localhost user=root database=testT`: failed SASL auth (FATAL: password authentication failed for user \"root\" (SQLSTATE 28P01))"))
	})
}
