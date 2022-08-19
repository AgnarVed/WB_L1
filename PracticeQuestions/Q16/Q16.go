package main

import (
	"fmt"
	"math/rand"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка

func main() {

	slice := []int{16, 213, 22, 575, 35, 78, -45, 3, 32}
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	quicksort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
