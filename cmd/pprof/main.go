package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigSystem := make(chan os.Signal, 1)

	signal.Notify(sigSystem, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// you can use this body to run pprof
	// run: `make run_pprof`
	// after run: `make pprof_mem` or `make pprof_cpu`

	<-sigSystem
}
