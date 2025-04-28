package main

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
)

var PORT int = 8080

type MetaData struct {
		Port int
		Title string
	}

func RenderTemplate(w http.ResponseWriter, fileName string, data any) {
	t, err := template.ParseFiles("Public/" + fileName + ".html")
	if err != nil {
		log.Println("[html/template]: ", err)
	}
	t.Execute(w, data)
}

func root(w http.ResponseWriter, r *http.Request) {
	data := MetaData{Port: PORT, Title: "index"}
	RenderTemplate(w, "index", data)
}

func main() {
    fmt.Printf("Starting HTTP server at localhost:%d\n", PORT)
	
	http.HandleFunc("/", root)
	
	port := fmt.Sprint(":", PORT) // Change the PORT variable from int to a string
	log.Fatal(http.ListenAndServe(port, nil))
}
