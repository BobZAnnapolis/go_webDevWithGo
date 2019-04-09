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

	fmt.Println("\nread ALL globs from folder, init slice, pass to template, print to screen, look at template to see how slice accessed\n")

	sagesSlice := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad", "Morrissey"}
	err := tpl.ExecuteTemplate(os.Stdout, "sliceToTemplate.gohtml", sagesSlice)
	if err != nil {
		log.Fatalln(err)
	}

}
