package main

import (
	"fmt"
	"os"
)

func main() {
	intCh := make(chan int)
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(sl); i++ {
		intCh <- sl[i]
	}
	close(intCh)

	go doSmth(intCh)

	for {
		num, opened := <-intCh
		if !opened {
			break
		}
		fmt.Fprintf(os.Stdout, "\nNumber is: %d", num)
	}

}

func doSmth(ch chan int) chan int {
	defer close(ch)
	result := make(chan int)
	var sl []int
	sl = append(sl, <-ch)
	for i := 0; i < len(sl); i++ {
		result <- sl[i] * 2
	}

	return result
}
