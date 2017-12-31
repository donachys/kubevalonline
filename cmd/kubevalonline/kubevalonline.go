package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/donachys/kubevalonline/api"
	"github.com/donachys/kubevalonline/app"
	"github.com/gorilla/handlers"
)

const (
	// Version of the application.
	Version = "0.0.0"
)

func main() {
	log.Println("Starting kubevalonline ", Version)
	Start()
}

func Start() {
	fs := flag.NewFlagSet("serve", flag.ExitOnError)
	httpAddr := fs.String("http", ":5000", "HTTP service address")
	m := http.NewServeMux()
	m.Handle("/api/", http.StripPrefix("/api", api.Handler()))
	m.Handle("/", app.Handler())

	log.Print("Listening on ", *httpAddr)
	httpServer := &http.Server{
		Addr:         *httpAddr,
		Handler:      handlers.CombinedLoggingHandler(os.Stdout, m),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatalln(httpServer.ListenAndServe())
}
