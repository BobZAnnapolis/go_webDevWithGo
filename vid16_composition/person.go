package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name string
	Age  int
}

func init() {
	tpl = template.Must(template.ParseGlob("html/*"))
}
func main() {

	fmt.Println("\nread ALL globs from folder, init struct, pass to template, print to screen, look at template to see it accessed\n")

	p1 := person{
		Name: "James Bond",
		Age:  42,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "person.gohtml", p1)
	if err != nil {
		log.Fatalln(err)
	}

}
