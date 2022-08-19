package main

import "fmt"

// Удалить i-ый элемент из слайса

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	newArr := deleteElement(3, arr)
	fmt.Println(newArr)
	arr[2] = 134
	fmt.Println(newArr)
}

func deleteElement(index int, arr []int) []int {
	res1 := arr[:index-1]
	res2 := arr[index:]
	res := append(res1, res2...)
	result := make([]int, len(arr)-1)
	copy(result, res)
	return result
}
