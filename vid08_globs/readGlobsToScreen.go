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

	fmt.Println("\nread ALL globs from folder, specify which one to write out to the screen\n")

	fmt.Println("\nwrite just glob3 to the screen\n")
	err := tpl.ExecuteTemplate(os.Stdout, "glob3.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\nwrite just glob1 to the screen\n")
	err = tpl.ExecuteTemplate(os.Stdout, "glob1.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\nwrite just glob2 to the screen\n")
	err = tpl.ExecuteTemplate(os.Stdout, "glob2.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\nwrite out the FIRST ONE IT FINDS to the screen\n")
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
