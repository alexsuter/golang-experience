package main

import (
	"fmt"
	"time"
)

func main() {
	chanApple := make(chan string)
	chanBanana := make(chan string)

	go func() {
		for {
			chanApple <- "Apple"
			time.Sleep(time.Second * 4)
		}
	}()
	go func() {
		for {
			chanBanana <- "Banana"
			time.Sleep(time.Second * 5)
		}
	}()

	for {
		select {
		case apple := <-chanApple:
			fmt.Println("Fruit: ", apple)
		case banana := <-chanBanana:
			fmt.Println("Fruit: ", banana)
		case <-time.After(time.Second * 2):
			fmt.Println("Ich habe 2 Sekunden gewartet und nichts erhalten.")
		}
	}
}
