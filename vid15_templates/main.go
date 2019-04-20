package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("html/*"))
}
func main() {

	fmt.Println("\nread ALL html and templates from folders, pass data to main file, display everything combined\n")

	err := tpl.ExecuteTemplate(os.Stdout, "displayTemplate.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}

}
