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

#import (
#    _ "net/http/pprof"
#    "net/http"
#)
#
#
#func main() {
#    go func() {
#        log.Println(http.ListenAndServe("localhost:6060", nil))
#    }()
#
#    // other code ...
#}
.PHONY: pprof_heap
pprof_heap:
	curl -o heap.out http://localhost:6060/debug/pprof/heap
	go tool pprof heap.out
