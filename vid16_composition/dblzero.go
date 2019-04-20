package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl1 *template.Template

type p2 struct {
	Name string
	Age  int
}

type doubleZero struct {
	p2
	LicenseToKill bool
}

func init() {
	tpl1 = template.Must(template.ParseGlob("html/*"))
}
func main() {

	fmt.Println("\nread ALL globs from folder, init struct, pass to template, print to screen, look at template to see it accessed\n")

	p1 := doubleZero{
		p2{
			Name: "James Bond",
			Age:  42,
		},
		false,
	}

	err := tpl1.ExecuteTemplate(os.Stdout, "dblzero.gohtml", p1)
	if err != nil {
		log.Fatalln(err)
	}

}
