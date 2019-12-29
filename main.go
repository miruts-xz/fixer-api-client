package main

import (
	"fixer-api-client/handler"
	"html/template"
	"net/http"
)

func main() {
	tmpl := template.Must(template.ParseGlob("template/*html"))
	requestHandler := handler.NewRequestHandler(tmpl)
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", requestHandler.Home)
	http.HandleFunc("/latest", requestHandler.Latest)
	http.HandleFunc("/historical", requestHandler.Historical)
	http.ListenAndServe(":8080", nil)
}
