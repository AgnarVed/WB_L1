package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

func main() {
	sl1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	sl2 := []int{2, 4, 6, 8, 10, 14, 23, 33, 222, 17}

	mp := make(map[int]bool)
	var result []int
	for _, val := range sl1 {
		mp[val] = true
	}
	for _, val := range sl2 {
		if ok := mp[val]; ok {
			result = append(result, val)
		}
	}
	fmt.Println(result)
}
