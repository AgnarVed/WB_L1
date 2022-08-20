package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow»

func main() {
	str := "snow dog sun buull daw awdjadw"
	fmt.Println(reverse(str))
	fmt.Println(reverse2(str))
}

func reverse(str string) string {
	nstr := strings.Split(str, " ")
	var res []string
	for i := len(nstr) - 1; i > -1; i-- {
		res = append(res, nstr[i])
	}
	return strings.Join(res, " ")
}

func reverse2(str string) string {
	words := strings.Fields(str)
	res := strings.Builder{}
	for i := len(words) - 1; i > -1; i-- {
		res.WriteString(words[i] + " ")
	}
	return res.String()
}
