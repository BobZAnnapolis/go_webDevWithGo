package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseGlob("html/*"))
}
func main() {

	fmt.Println("\nread ALL globs from folder, init slice of Structs, pass to template, print to screen, look at template to see how slice accessed\n")

	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	mlk := sage{
		Name:  "MLK",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	muhammad := sage{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil.",
	}

	morrissey := sage{
		Name:  "Morrissey",
		Motto: "Everyday is like Sunday",
	}

	sages := []sage{buddha, gandhi, mlk, jesus, muhammad, morrissey}

	err := tpl.ExecuteTemplate(os.Stdout, "sliceStructToTemplate.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}

}
