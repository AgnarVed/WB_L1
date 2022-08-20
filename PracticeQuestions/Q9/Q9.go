package main

import (
	"fmt"
	"sync"
)

// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.

func main() {
	inCh := make(chan int)
	outCh := make(chan int)
	//sl := []int64{2, 4, 6, 8, 10}
	sl := make([]int, 100)
	for i := range sl {
		sl[i] = i + 1
	}
	wg := sync.WaitGroup{}
	wg.Add(3)
	go go1(inCh, sl, &wg)
	go multy(inCh, outCh, &wg)
	go printResult(outCh, &wg)
	wg.Wait()
	fmt.Println("Finished")

}

func go1(ch chan int, sl []int, wg *sync.WaitGroup) {
	defer close(ch)
	for _, val := range sl {
		ch <- val
	}
	wg.Done()
}

func multy(ch, ch2 chan int, wg *sync.WaitGroup) {
	defer close(ch2)
	for val := range ch {
		ch2 <- val * 2
	}
	wg.Done()
}

func printResult(ch chan int, wg *sync.WaitGroup) {
	for val := range ch {
		fmt.Println("Result is: ", val)
	}
	wg.Done()
}
