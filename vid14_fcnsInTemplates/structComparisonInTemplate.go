package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl3 *template.Template

func init() {
	tpl3 = template.Must(template.ParseFiles("html/structComparisonInTemplate.gohtml"))
}
func main() {

	fmt.Println("\nread template from folder, init struct, pass into template, print using fcns in template to screen, look at template to see how fcn works\n")

	g1 := struct {
		Score1 int
		Score2 int
	}{
		7,
		9,
	}

	err := tpl3.Execute(os.Stdout, g1)
	if err != nil {
		log.Fatalln(err)
	}

}
