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

	fmt.Println("\nread ALL globs from folder, init map, pass to template, print to screen, look at template to see it accessed\n")

	sagesMap := map[string]string{
		"India":    "Gandhi",
		"USA":      "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Muhammad",
		"World":    "Morrissey",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "mapToTemplate.gohtml", sagesMap)
	if err != nil {
		log.Fatalln(err)
	}

}
