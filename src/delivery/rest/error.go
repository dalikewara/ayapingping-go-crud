package rest

import (
	"github.com/dalikewara/ayapingping-go-crud/src/entity"
	"net/http"
)

// Parse

var ErrParseJSONRequest = entity.NewStdError("REST-PARSE-01", "error parsing json request", http.StatusBadRequest)

// Compose

var ErrComposeResponseData = entity.NewStdError("REST-COMPOSE-01", "error composing response data", http.StatusInternalServerError)
