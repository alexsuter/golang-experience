package main

import (
  "fmt"
  "time"
)

func main() {
    workIsDone := make(chan bool)

	go doSomeWork(workIsDone)

	fmt.Println("Warten auf das Ende von doSomeWork")
	<-workIsDone
	fmt.Println("doSomeWork ist beendet")
}

func doSomeWork(workIsDone chan bool) {
	fmt.Println("do work 1")
	time.Sleep(2 * time.Second)
	fmt.Println("do work 2")
	workIsDone <- true
}