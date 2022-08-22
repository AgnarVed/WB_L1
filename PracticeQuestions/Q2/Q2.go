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
	wg.Add(1)
	Sqr(arr)
	go SqrConc(arr, &wg)

	wg.Wait()

}

func Sqr(arr []int) {
	for _, val := range arr {
		res := val * val
		fmt.Fprintln(os.Stdout, res)
	}
}

func SqrConc(arr []int, wg *sync.WaitGroup) {
	wg.Add(len(arr) - 1)
	for i := 0; i < len(arr); i++ {
		go func(wg *sync.WaitGroup, val int) {
			res := val * val
			fmt.Fprintln(os.Stdout, res)
			wg.Done()
		}(wg, arr[i])
	}
}
