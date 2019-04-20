package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl1 *template.Template

func init() {
	tpl1 = template.Must(template.ParseFiles("html/structFcnInTemplate.gohtml"))
}
func main() {

	fmt.Println("\nread template from folder, init struct, pass into template, print using fcns in template to screen, look at template to see how fcn works\n")

	xs := []string{"zero", "one", "two", "three", "four", "five"}

	data := struct {
		Words []string
		Lname string
	}{
		xs,
		"Zawislak",
	}

	err := tpl1.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

}
