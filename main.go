package main

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	mutex sync.Mutex
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	path := r.URL.Path
	fmt.Fprintf(w, "%s", path)
}

func main() {
	fmt.Println("Starting The Web Server on port 8000")
	http.HandleFunc("/", HandleRequest)
	server := &http.Server{
		Addr: "localhost:8000",
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("Error with Starting Server", err)
	}
}
