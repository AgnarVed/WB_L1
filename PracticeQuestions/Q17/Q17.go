package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(63, items))
}

func binarySearch(value int, input []int) bool {

	left := 0
	right := len(input) - 1

	for left <= right {
		middle := (left + right) / 2

		if input[middle] < value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	if left == len(input) || input[left] != value {
		return false
	}

	return true
}
