info:
	@echo "Makefile is your friend"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

start-rest: ## start rest app
	@go run src/app/rest/rest.go

mock: ## generates mocks
	@go install github.com/vektra/mockery/v2@latest
	@mockery --dir=src/repository --all --output=src/repository/mocks
	@mockery --dir=src/service --all --output=src/service/mocks

integration-test: ## runs integration tests
	@- go test ./... -v -cover -tags=integration > integration_test.out
	@cat integration_test.out

unit-test: ## runs unit tests
	@- go test ./... -v -cover -tags=unit > unit_test.out
	@cat unit_test.out