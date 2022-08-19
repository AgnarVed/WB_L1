package main

import (
	"fmt"
	"os"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов чисел
//взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout

func main() {
	arr := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	for _, val := range arr {
		go func(val int) {
			res := val * val
			fmt.Fprintln(os.Stdout, res)
			defer wg.Done()
		}(val)
	}
	wg.Wait()

}
