package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"text/template"
)

var tpl2 *template.Template

// create a FuncMap to register functions
// "uc" is what the func will be called in the template
// "uc" is the ToUpper function from package strings
// "ft" is a function I declare
// "ft" slices a string, returning the 1st 3 characters
var fm2 = template.FuncMap{
	"fdbl":  double,
	"fsq":   square,
	"fsqrt": sqroot,
}

func init() {
	tpl2 = template.Must(template.New("").Funcs(fm2).ParseFiles("html/pipelining.gohtml"))
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqroot(x float64) float64 {
	return math.Sqrt(x)
}

func main() {

	fmt.Println("\nread html file from folder, define funcmap, call template w/val 3 to be use fcns to square and square root\n")

	err := tpl2.ExecuteTemplate(os.Stdout, "pipelining.gohtml", 3)
	if err != nil {
		log.Fatalln(err)
	}

}
