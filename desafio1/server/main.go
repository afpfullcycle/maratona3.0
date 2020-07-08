package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type TitleHelloFC struct {
	HelloFCTitle string
}

func main()  {
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/hello", HelloFCHandler)
	http.ListenAndServe(":8080", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Full Cycle")
}

func HelloFCHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("hello.html"))
	hello := TitleHelloFC{HelloFCTitle: "Hello Full Cycle"}
	tmpl.Execute(w, hello)
}
