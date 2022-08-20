package main

import (
	"fmt"
	"os"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов чисел
//взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout

func main() {
	arr := []int{1, 2, 3, 4, 5}
	wg := sync.WaitGroup{}
	wg.Add(2)
	Sqr(arr)
	SqrConc(arr, &wg)
	SqrConc2(arr, &wg)

	wg.Wait()

}

func Sqr(arr []int) {
	for _, val := range arr {
		res := val * val
		fmt.Fprintln(os.Stdout, res)
	}
}

func SqrConc(arr []int, wg *sync.WaitGroup) {
	go func(wg *sync.WaitGroup) {
		for _, val := range arr {
			res := val * val
			fmt.Fprintln(os.Stdout, res)
		}
		defer wg.Done()
	}(wg)
}

func SqrConc2(arr []int, wg *sync.WaitGroup) {
	wg.Add(len(arr) - 1)
	for _, val := range arr {
		go func(wg *sync.WaitGroup, val int) {
			res := val * val
			fmt.Fprintln(os.Stdout, res)
			wg.Done()
		}(wg, val)
	}
}
