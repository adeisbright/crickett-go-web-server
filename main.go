package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

var (
	mutex sync.Mutex
	// activeReqs int
	// wg         sync.WaitGroup
)

// func logRequest(r *http.Request) {
// 	defer wg.Done()
// 	mutex.Lock()
// 	activeReqs++
// 	fmt.Printf("Received request for: %s. Active requests: %d\n", r.URL.Path, activeReqs)
// 	mutex.Unlock()
// }

type templateData = interface{}

func renderTemplate(templatePath string, w http.ResponseWriter, data templateData) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		notFoundTemplate, _ := template.ParseFiles("template/" + "404" + ".html")
		notFoundTemplate.Execute(w, data)
		return
	}

	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

type User struct {
	Name       string
	IsEngineer bool
	Skills     []string
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	path := r.URL.Path
	templatePath := ""
	if path == "/" {
		templatePath = "template/" + "index" + ".html"
	} else {
		templatePath = "template" + path + ".html"
	}
	user := &User{
		Name:       "Adeleke Bright",
		IsEngineer: true,
		Skills:     []string{"Go", "Typescript", "Python", ".NET"},
	}
	renderTemplate(templatePath, w, user)

	// t, err := template.ParseFiles(templatePath)
	// if err != nil {
	// 	http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// if err := t.Execute(w, nil); err != nil {
	// 	http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// wg.Add(1)
	// go func() {
	// 	logRequest(r)

	// 	t, err := template.ParseFiles("template/index.html")
	// 	if err != nil {
	// 		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	if err := t.Execute(w, nil); err != nil {
	// 		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// }()
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

	// wg.Add(1)
	// go func() {
	// 	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		fmt.Println("Error with Starting Server", err)
	// 	}
	// }()

	// wg.Wait()
}
