package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	sl := []string{"cat", "cat", "dog", "cat", "tree", "man", "wild berry", "donut", "donut", "wild berry"}
	fmt.Println(getSet(sl))
}

func getSet(sl []string) []string {
	var mmbr struct{}
	var result []string
	set := make(map[interface{}]struct{})
	for _, val := range sl {
		set[val] = mmbr
	}
	for k := range set {
		result = append(result, k.(string))
	}
	return result
}
