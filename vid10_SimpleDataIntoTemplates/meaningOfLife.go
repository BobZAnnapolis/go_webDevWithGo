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

	fmt.Println("\nread ALL globs from folder, pass data {ie 42} to meaningOfLife file, output combined template & data to screen\n")

	err := tpl.ExecuteTemplate(os.Stdout, "meaningOfLife.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}

}
