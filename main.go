package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("pong")
}
func statusHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	mutex.Unlock()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Fprintf(w, "Visistas: %d\nCorriendo en el puerto %s\n", counter, port)
}

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/status", statusHandler)

	log.Println("Servidor corriendo en el puerto", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
