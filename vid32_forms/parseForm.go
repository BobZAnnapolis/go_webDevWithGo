package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("html/*"))
}

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "parseForm.html", req.Form)
}

func main() {

	fmt.Println("\nread ALL htmls from folder, init struct, pass to template, look at template to see it accessed\n")

	var d hotdog
	http.ListenAndServe(":8080", d)

}
