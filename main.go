package main

import (
    "html/template"
    "net/http"
)

var counter int

func indexHandler(w http.ResponseWriter, r *http.Request) {
	counter = 0
    t, _ := template.ParseFiles("templates/index.html")
    t.Execute(w, counter)
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
    counter++
    t, _ := template.New("counter").Parse("{{ . }}")
    t.Execute(w, counter)
}

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/increment", incrementCounter)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    http.ListenAndServe(":8080", nil)
}

