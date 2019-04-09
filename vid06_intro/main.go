package main

import (
	"fmt"
)

var xPkgLvl = 13

type person struct {
	fname string
	lname string
	age   int
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, " says, `Person struct says Good morning, BJ\n\n")
}

func (s secretAgent) speak() {
	fmt.Println(s.fname, s.lname, " says, `Secret Agent say Good morning, BJ\n\n")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	x := 7
	fmt.Printf("\nshort variable declaration x:=7\nfmt.Println(x) : %v\n", x)

	fmt.Printf("\nPrint the type of variable using percent T, type : %t\n", x)

	fmt.Printf("\nDeclare above func main for package level scope var xPkgLvl=13, inside fcn print xPkgLvl : %v\n", xPkgLvl)

	fmt.Println("\n\nComposite Literals\n")
	fmt.Println("\nSlices - hold a list of things, ints, strings, etc\n\txi := []int{2,4,7,9,42}\n\tfmt.Println(xi)\n\t")
	xi := []int{2, 4, 7, 9, 42}
	fmt.Println(xi)

	fmt.Println("\n\nMaps, ie key-value pairs, quick lookup - \n\tmap := map[string]int{\n\t\t'Todd': 45,\n\t\t'Job': 42,\n\t\t}\n\t\tfmt.Println(map)")
	m := map[string]int{
		"Todd": 45,
		"Job":  42,
	}
	fmt.Println(m)

	fmt.Println("\n\nStructs, ie user-defined app-specific - \n\ttype person struct {\n\t\tfname string\n\t\tlname string\n\t\tage int\n\t}\np1 := person {\n\t'Miss',\n\t'Moneypenny'.\n}\nfmt.Println(p1)\n")
	p1 := person{
		"Miss",
		"MoneyPenny",
		13,
	}
	fmt.Println(p1)
	p1.speak()

	sa1 := secretAgent{
		p1,
		false,
	}
	sa2 := secretAgent{
		person{
			"James",
			"Bond",
			41,
		},
		true,
	}
	fmt.Println(sa1)
	fmt.Println(sa2)

	sa1.speak()
	sa2.speak()
	sa2.person.speak()

	saySomething(p1)
	saySomething(sa1)
	saySomething(sa2)
}
