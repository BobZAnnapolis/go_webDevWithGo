package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type Page struct {
	Title   string
	Heading string
	Body    string
}

func init() {
	tpl = template.Must(template.ParseGlob("html/*"))
}
func main() {

	fmt.Println("\nread ALL htmls from folder, init struct, pass to template, look at template to see it accessed\n")

	home := Page{
		Title:   "Nothing Escaped",
		Heading: "Nothing is escaped in text/template",
		Body:    `<script>alert("Yow ! - javascript insertion")</script>`,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "xss.gohtml", home)
	if err != nil {
		log.Fatalln(err)
	}

}
