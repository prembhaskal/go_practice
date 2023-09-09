package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var port int

func main() {
	flag.IntVar(&port, "port", 8090, "port where server to start on")
	flag.Parse()

	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		Handler: http.HandlerFunc(returnTimeHandler),
	}
	err := server.ListenAndServe()
	fmt.Printf("server closed: %v", err)
}

func returnTimeHandler(w http.ResponseWriter, r *http.Request) {
	currtime := time.Now()
	currtimestr := currtime.Format(time.RFC3339Nano)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "time now is :%s\n", currtimestr)
}
