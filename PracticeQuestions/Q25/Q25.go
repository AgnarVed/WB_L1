package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep

func main() {
	//customSleep(1)
	//var input string
	//fmt.Scanln(input)
	Sleep(2)
	fmt.Println("Wake up!")
}

func Sleep(x int) {
	<-time.After(time.Duration(x) * time.Second)
}

func customSleep(duration int) {
	ch := time.After(time.Second * time.Duration(duration))
	for {
		select {
		case <-ch:
			fmt.Println("slept!")
			return // MUST RETURN, else endless loop!
		default:
			fmt.Println("Waiting")
		}
	}
}
