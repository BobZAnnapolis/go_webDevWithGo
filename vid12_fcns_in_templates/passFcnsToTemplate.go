package main

import (
	"fmt"
	"log"
	"os"
	"strings"
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

// create a FuncMap to register functions
// "uc" is what the func will be called in the template
// "uc" is the ToUpper function from package strings
// "ft" is a function I declare
// "ft" slices a string, returning the 1st 3 characters
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("html/passFcnsToTemplate.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {

	fmt.Println("\nread ALL globs from folder, init struct slice struct, pass to template, print to screen seeing how fcns operate inside template\n")

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
		Model:        "Mazda3",
		Doors:        2,
	}

	sages := []sage{m, n}
	cars := []car{f, t, z}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "passFcnsToTemplate.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}
