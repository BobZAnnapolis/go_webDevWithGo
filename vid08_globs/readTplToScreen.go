package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {

	fmt.Println("\nread text from 1 templated file, write to the screen\n")

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\nadd tpl2 & tpl3 templated files into container\n")
	tpl, err = tpl.ParseFiles("tpl2.gohtml", "tpl3.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\nwrite just tpl3 to the screen\n")
	err = tpl.ExecuteTemplate(os.Stdout, "tpl3.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\nwrite just tpl2 to the screen\n")
	err = tpl.ExecuteTemplate(os.Stdout, "tpl2.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\nwrite out the FIRST ONE IT FINDS to the screen\n")
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
