package main

import (
  "fmt"
  "time"
)

func main() {
    fruits := make(chan string, 2)

    go func() {
      fruits <- "Apple"
      fmt.Println("Apple im Channel")
      fruits <- "Banana"
      fmt.Println("Banana im Channel")
      fruits <- "Orange"
      fmt.Println("Orange im Channel")
    } ()

    time.Sleep(3 * time.Second)

    fruit := <-fruits
    fmt.Println(fruit)

    time.Sleep(3 * time.Second)

    fruit = <-fruits
    fmt.Println(fruit)

    time.Sleep(3 * time.Second)

    fruit = <-fruits
    fmt.Println(fruit)
}
