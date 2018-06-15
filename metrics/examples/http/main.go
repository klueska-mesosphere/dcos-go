package main

import (
	"fmt"
	"net/http"
	"time"

	. "github.com/dcos/dcos-go/dcos-metrics-generator"
	. "github.com/dcos/dcos-go/dcos-metrics-generator/http/middleware"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	handler := http.HandlerFunc(simpleHandler)

	scope, closer := NewDCOSComponentScope("dcos-log", time.Second, nil, nil, true)
	defer closer.Close()

	router.Handle("/", StatsMiddleware(handler, scope))
	http.ListenAndServe("127.0.0.1:8080", router)
}

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}
