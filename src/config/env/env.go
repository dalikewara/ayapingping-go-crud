package env

import (
	"github.com/dalikewara/ayapingping-go-crud/src/config/constant"
	"github.com/joho/godotenv"
	"os"
)

var _ = godotenv.Load(constant.EnvFilePath)

var AppEnv = os.Getenv("APP_ENV")

var TimezoneName = os.Getenv("TIMEZONE_NAME")
var TimezoneOffset = os.Getenv("TIMEZONE_OFFSET")

var RESTPort = os.Getenv("REST_PORT")

var PostgreSQLHost = os.Getenv("POSTGRESQL_HOST")
var PostgreSQLPort = os.Getenv("POSTGRESQL_PORT")
var PostgreSQLUser = os.Getenv("POSTGRESQL_USER")
var PostgreSQLPass = os.Getenv("POSTGRESQL_PASS")
var PostgreSQLDBName = os.Getenv("POSTGRESQL_DBNAME")
var PostgreSQLOption = os.Getenv("POSTGRESQL_OPTION")
