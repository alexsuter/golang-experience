package main

import (
	"fmt"
	"strings"
	"strconv"
	)
 
func main() {
	
	anna := Person{name: "Anna", tel: 147}
	peter := Person{name: "Peter", tel: 581}
	
	book := Book{personen: []Person{anna, peter}} 
	
	fmt.Println("Suchen nach Anna")
	tel, found := book.lookAt("Anna")
	if found {
		fmt.Println("Number of Anna is " + strconv.Itoa(tel))	
	}
	
	fmt.Println("Suchen nach Marc")
	tel, found = book.lookAt("Marc")
	if !found {
		fmt.Println("Number of Marc not found")	
	}
	
	
}

type Book struct {
	personen []Person
}
 
type Person struct {
  name string
  tel int
}

func (book Book) lookAt(name string) (int, bool) {
    for _, p := range book.personen {
    	if strings.EqualFold(p.name, name) {
    		return p.tel, true
    	}
    }
    return 0, false
}