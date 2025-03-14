default: help

.PHONY: run
run: ## running a telegram bot via linux
	go run ./cmd/tg_bot/main.go

.PHONY: run_print_json
run_print_json: ## running a json printer via linux
	go run ./cmd/print_json/main.go

.PHONY: test
test: ## running a test via linux
	go test -v -coverpkg=./internal/...,./pkg/... ./...

.PHONY: test_coverage
test_coverage: ## running a test coverage via linux
	go test -coverprofile=coverage.out -coverpkg=./internal/...,./pkg/... ./...
	go tool cover -html=coverage.out

.PHONY: mockgen
mockgen: ## install mockgen 'go install go.uber.org/mock/mockgen@latest'
	go generate ./...

.PHONY: lint
lint: ## run golangci-lint (2-minute wait)
	docker run -t --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.64 golangci-lint run -v

.PHONY: run_pprof
run_pprof: ## run after 'make pprof_mem' or 'make pprof_cpu'
	go run ./cmd/pprof/main.go

.PHONY: pprof_mem
pprof_mem: ## run before 'make run_pprof'. analyze memory
	curl -o mem.prof http://localhost:6060/debug/pprof/heap
	go tool pprof -http=:8080 mem.prof

.PHONY: pprof_cpu
pprof_cpu: ## run before 'make run_pprof'. analyze cpu
	curl -o cpu.prof http://localhost:6060/debug/pprof/profile
	go tool pprof -http=:8080 cpu.prof

.PHONY: build
build: ## build app for linux
	CGO_ENABLED=0 GOOS=linux go build -o tg_reaction ./cmd/tg_bot/main.go

.PHONY: docker_build
docker_build: ## build app for docker
	docker build . -t ghcr.io/fromsi/tg_reaction:latest

APP_TELEGRAM_TOKEN ?= # Telegram bot token

.PHONY: docker_run_img
docker_run_img: ## run telegram bot via docker. APP_TELEGRAM_TOKEN is required
	docker run --rm -e APP_TELEGRAM_TOKEN=$(APP_TELEGRAM_TOKEN) ghcr.io/fromsi/tg_reaction:latest

.PHONY: docker_push
docker_push: ## open link https://github.com/settings/tokens > Generate New Token > Classic > write:packages
	echo "For login to ghcr.io use: docker login ghcr.io -u USERNAME -p TOKEN"
	docker push ghcr.io/fromsi/tg_reaction:latest

.PHONY: help
help: ## display this help message
	@cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
