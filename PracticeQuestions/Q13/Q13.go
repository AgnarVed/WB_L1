package main

import "fmt"

//Поменять местами два числа без создания временной переменной (too easy?)

func main() {
	a := 2
	b := 3
	a, b = b, a
	fmt.Println(a, b)
}
