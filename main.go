package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	cfg config
	log *log.Logger
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API srever port")
	flag.StringVar(&cfg.env, "enviornment", "development", "Enviornment(development|staging|production)")

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		cfg: cfg,
		log: logger,
	}

	// mux := http.NewServeMux()
	// mux.HandleFunc("/v1/healthCheck", app.healthCheckHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
