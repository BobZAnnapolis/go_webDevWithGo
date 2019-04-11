package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

var tpl1 *template.Template

// create a FuncMap to register functions
// "uc" is what the func will be called in the template
// "uc" is the ToUpper function from package strings
// "ft" is a function I declare
// "ft" slices a string, returning the 1st 3 characters
var fm1 = template.FuncMap{
	"fdateMDY": monthDayYear,
	"kitchen":  kitchen,
}

func init() {
	tpl1 = template.Must(template.New("").Funcs(fm1).ParseFiles("html/timeFormatting.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func kitchen(t time.Time) string {
	return t.Format(time.Kitchen)
}

func main() {

	fmt.Println("\nread html file from folder, define funcmap, call template w/time formatting fcn\n")

	err := tpl1.ExecuteTemplate(os.Stdout, "timeFormatting.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}

}
