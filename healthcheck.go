package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: avialable")
	fmt.Fprintf(w, "enviornment: %s\n", app.cfg.env)
	fmt.Fprintf(w, "port: %d\n", app.cfg.port)
}
