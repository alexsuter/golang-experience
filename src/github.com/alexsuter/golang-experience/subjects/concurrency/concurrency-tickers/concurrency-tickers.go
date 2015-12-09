package main

import (
	"fmt"
	"time"
)

func main() {
	fruitTicker := time.NewTicker(time.Second)
    go func() {
        for t := range fruitTicker.C {
            fmt.Println("Fruit at", t)
        }
    }()
    time.Sleep(time.Second * 6)
    fruitTicker.Stop()
}
