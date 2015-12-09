package main

import (
  "fmt"
)

func main() {
    
    fruits := make(chan string)
    done := make(chan bool)
    
    go func() {
    	for {
    		fruit, open := <-fruits
    		if (open) {
    			fmt.Println("Fruit: ", fruit)
    		} else {
    			fmt.Println("no more fruits")
    			done <- true
    			return
    		}
    	}
    }()
    
    fruits <- "Apple"
    fruits <- "Banana"
    fruits <- "Orange"
    
    close(fruits)
    
    // fruits <- "Orange2" Runtime Error
    
    <-done
    
    // Range over channel
    cars := make(chan string, 2)
    cars <- "Audi"
    cars <- "VW"
    //close(cars)
    for car := range cars {
    	fmt.Println(car)
    }
}


