package main

import (
	"fmt"
	"strings"
	"strconv"
	)
 
func main() {
	book := Book{}

	book.add("Anna", 147)
	book.add("Peter", 581)
	
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

func (book *Book) add(name string, tel int) {
	for _, p := range book.personen {
    	fmt.Println(p.name)	
    }
	book.personen = append(book.personen, Person{name: name, tel: tel})
	for _, p := range book.personen {
    	fmt.Println(p.name)	
    }
}

func (book *Book) lookAt(name string) (int, bool) {
    for _, p := range book.personen {
    	if strings.EqualFold(p.name, name) {
    		return p.tel, true
    	}
    }
    
    return 0, false
}