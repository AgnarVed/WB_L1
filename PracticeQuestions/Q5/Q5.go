package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

func main() {
	var duration int // время выполнения
	duration = 5
	msgChan := make(chan interface{}) // канал
	defer close(msgChan)              // закрытие канала

	finishTime := time.Duration(duration) * time.Second // время окончания

	startTime := time.Now() // время начала

	// горутина, что выводит данные из канала
	go func(ch chan interface{}) {
		for val := range ch {
			fmt.Println("Value is: ", val)
		}
	}(msgChan)

	// генератор случайных значений
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	for time.Since(startTime) < finishTime {
		val := generator.Intn(100)
		msgChan <- val // отправка значений в канал
		time.Sleep(500 * time.Millisecond)
	}
}
