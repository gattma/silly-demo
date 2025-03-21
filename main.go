package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Silly demo!")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	bgColor := os.Getenv("BG_COLOR")
	if bgColor == "" {
		bgColor = "black"
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		BgColor string
	}{
		BgColor: bgColor,
	}
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/", indexHandler)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
