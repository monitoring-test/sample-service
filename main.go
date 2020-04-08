package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Print("Started service")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)

	// HTTP requests handle loop
	go func() {
		if err := http.ListenAndServe(":4569", r); err != nil {
			log.Fatalf("buggy bug %q", err.Error())
		}
	}()

	// Wait for stop
	quit := make(chan os.Signal, 5)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	<-quit
	log.Print("Shut down app")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	help := "This is sample service (WIP)"
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte(help))
}
