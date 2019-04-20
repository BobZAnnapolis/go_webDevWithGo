package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl2 *template.Template

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

func init() {
	tpl2 = template.Must(template.ParseGlob("html/*"))
}
func main() {

	fmt.Println("\nread ALL globs from folder, init struct, pass to template, print to screen, look at template to see it accessed\n")

	p1 := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-40", "Intro to Programming In Go", "4"},
				course{"CSCI-130", "Intro to Web Programming with Go", "4"},
				course{"CSCI-140", "Mobile Apps Using Go", "4"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"CSCI-50", "Advanced Go", "5"},
				course{"CSCI-190", "Advanced Web Programming with Go", "5"},
				course{"CSCI-191", "Advanced Mobile Apps Using Go", "5"},
			},
		},
	}
	err := tpl2.ExecuteTemplate(os.Stdout, "courses.gohtml", p1)
	if err != nil {
		log.Fatalln(err)
	}

}
