package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl2 *template.Template

type usersStruct struct {
	Name  string
	Motto string
	Admin bool
}

func init() {
	tpl2 = template.Must(template.ParseFiles("html/arrayStructsToTemplate.gohtml"))
}
func main() {

	fmt.Println("\nread file from folder, init data struct, pass to template, print to screen, look at template to see how data accessed\n")

	u1 := usersStruct{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
		Admin: false,
	}

	u2 := usersStruct{
		Name:  "Gandhi",
		Motto: "Be the change",
		Admin: true,
	}

	u3 := usersStruct{
		Name:  "",
		Motto: "The man is nothing.",
		Admin: true,
	}

	users := []usersStruct{u1, u2, u3}

	err := tpl2.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}

}
