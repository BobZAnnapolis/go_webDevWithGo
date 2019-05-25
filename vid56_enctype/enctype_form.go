package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("html/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {

	fmt.Println("\nenctype=multipart form\nvisit this page : localhost:8080")

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	// body
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	body := string(bs)

	err := tpl.ExecuteTemplate(w, "enctype_form.html", body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
