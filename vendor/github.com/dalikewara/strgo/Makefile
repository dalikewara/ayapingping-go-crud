info:
	@echo "Makefile is your friend"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

test: ## run test cases
	@- go test -cover ./... -v > test.out
	@cat test.out

test-benchmark: ## run benchmark
	@go test -bench=. -benchmem > benchmark.out
	@cat benchmark.out