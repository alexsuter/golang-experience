package main

import (
  "fmt"
  "time"
)

func main() {
    printMe("Synchroner Peter", 5)

    go printMe("Asynchroner Markus", 5)
    go printMe("Asynchrone Anna", 5)

    printMe("Synchhoner Fabian", 5)

    var input string
    fmt.Scanln(&input)
}

func printMe(text string, times int) {
    for i := 0; i < times; i++ {
        time.Sleep(1000)
        fmt.Println(text, ":", i)
    }
}
