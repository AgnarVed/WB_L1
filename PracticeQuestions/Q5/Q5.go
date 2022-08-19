package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

func main() {
	var duration int
	duration = 5
	msgChan := make(chan interface{})
	defer close(msgChan)

	finishTime := time.Duration(duration) * time.Second

	startTime := time.Now()

	go func(ch chan interface{}) {
		for val := range ch {
			fmt.Println("Value is: ", val)
		}
	}(msgChan)

	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	for time.Since(startTime) < finishTime {
		val := generator.Intn(100)
		msgChan <- val
		time.Sleep(500 * time.Millisecond)
	}
}
