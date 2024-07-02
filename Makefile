.PHONY: run
run:
	go run main.go

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -o tg_reaction .

.PHONY: docker_build
docker_build:
	docker build . -t ghcr.io/fromsi/tg_reaction:latest

.PHONY: docker_push
docker_push:
	docker push ghcr.io/fromsi/tg_reaction:latest
