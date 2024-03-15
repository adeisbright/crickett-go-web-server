package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type templateData = interface{}

func renderTemplate(templatePath string, w http.ResponseWriter, data templateData) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		notFoundTemplate, _ := template.ParseFiles("template/" + "404" + ".html")
		notFoundTemplate.Execute(w, nil)
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
