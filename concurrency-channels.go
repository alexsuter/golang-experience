package main

import (
  "fmt"
  "time"
)

func main() {
    fruits := make(chan string)

    go func() {
      fruits <- "Apple"
      time.Sleep(3 * time.Second)
      fruits <- "Banana"
      fruits <- "Orange"
    } ()

    fruit := <-fruits
    fmt.Println(fruit)

    fruit = <-fruits
    fmt.Println(fruit)

    time.Sleep(3 * time.Second)

    fruit = <-fruits
    fmt.Println(fruit)
}
