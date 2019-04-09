package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sagesStruct struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseGlob("html/*"))
}
func main() {

	fmt.Println("\nread ALL globs from folder, init struct, pass to template, print to screen, look at template to see how slice accessed\n")
	sages := sagesStruct{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "structToTemplate.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}

}
