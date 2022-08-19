package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет,
//что все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.

func main() {
	fmt.Println(unique("abcd"))
	fmt.Println(unique("abCdefAaf"))
	fmt.Println(unique("aabcd"))
}

func unique(input string) bool {
	lowerCase := strings.ToLower(input)
	tmp := make(map[rune]bool)

	for _, val := range lowerCase {
		if ok := tmp[val]; !ok {
			tmp[val] = true
		} else {
			return false
		}
	}
	return true
}
