package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("html/sliceFcnInTemplate.gohtml"))
}
func main() {

	fmt.Println("\nread template from folder, init slice, pass into template, print using slice fcn in template to screen, look at template to see how fcn works\n")

	xs := []string{"zero", "one", "two", "three", "four", "five"}
	err := tpl.Execute(os.Stdout, xs)
	if err != nil {
		log.Fatalln(err)
	}

}
