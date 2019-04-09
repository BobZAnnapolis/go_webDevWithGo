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

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseGlob("html/*"))
}
func main() {

	fmt.Println("\nread ALL globs from folder, init struct slice struct, pass to template, print to screen, look at template to see how slice accessed\n")

	m := sage{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil.",
	}

	n := sage{
		Name:  "Morrissey",
		Motto: "Everyday is like Sunday",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	t := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	z := car{
		Manufacturer: "Mazda",
		Model:        "3",
		Doors:        2,
	}

	sages := []sage{m, n}
	cars := []car{f, t, z}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "structSliceStructToTemplate.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}
