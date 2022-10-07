package constant

import (
	"net/http"
)

const EnvFilePath = ".env"

const AppEnvProduction = "production"

const DefaultOkCode = "00"
const DefaultOkMessage = "ok"
const DefaultOkStatus = http.StatusOK
const DefaultNotOkStatus = http.StatusInternalServerError
