package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("html/parseForm1.html"))
}

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}

	tpl.ExecuteTemplate(w, "parseForm1.html", data)
}

func main() {

	fmt.Println("\nread ALL htmls from folder, init struct, pass to template, look at template to see it accessed\n")

	var d hotdog
	http.ListenAndServe(":8080", d)

}
