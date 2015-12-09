package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 2)
	<-timer.C
	fmt.Println("2 Sekunden gewartet")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Time abgelaufen")
		
	}()
	timer2.Stop()
}
