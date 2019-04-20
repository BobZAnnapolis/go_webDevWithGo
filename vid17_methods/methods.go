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

func (p person) SomeProcessing() int {
	return 7
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) TakesArgs(x int) int {
	return x * 2
}

func init() {
	tpl = template.Must(template.ParseGlob("html/*"))
}
func main() {

	fmt.Println("\nread ALL globs from folder, init struct, pass to template, print to screen, look at template to see it accessed\n")

	p1 := person{
		Name: "Ian Fleming",
		Age:  56,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "methods.gohtml", p1)
	if err != nil {
		log.Fatalln(err)
	}

}
