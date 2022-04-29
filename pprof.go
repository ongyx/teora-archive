//go:build pprof

package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func pprof() {
	log.Println(http.ListenAndServe("localhost:6060", nil))
}
