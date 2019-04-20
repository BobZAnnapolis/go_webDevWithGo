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

	fmt.Println("\nread ALL globs from folder, pass string into varInTemplate file, output combined template & data to screen\n")

	err := tpl.ExecuteTemplate(os.Stdout, "varInTemplate.gohtml", "Release self-focus; embrace other-focus")
	if err != nil {
		log.Fatalln(err)
	}

}
