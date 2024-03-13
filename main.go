package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

var (
	mutex sync.Mutex
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	t, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
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
