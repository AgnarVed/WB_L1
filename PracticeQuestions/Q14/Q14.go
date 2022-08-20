package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}

func main() {
	input := make([]interface{}, 5)
	var n int
	var s string
	var b bool
	c := make(chan int)
	c2 := make(chan struct{})
	input[0] = n
	input[1] = s
	input[2] = b
	input[3] = c
	input[4] = c2
	//fmt.Printf("Type of variable is: %T\n", input[0].(int))
	//fmt.Printf("Type of variable is: %T\n", input[1].(string))
	//fmt.Printf("Type of variable is: %T\n", input[2].(bool))
	//fmt.Printf("Type of variable is: %T\n", input[3].(chan int))
	//fmt.Printf("Type of variable is: %T\n", input[4].(chan struct{}))
	fmt.Println(CheckType(n))
	fmt.Println(CheckType(s))
	fmt.Println(CheckType(b))
	fmt.Println(CheckType(c))
	fmt.Println(CheckType(c2))

	fmt.Println(reflectType(n))
	fmt.Println(reflectType(s))
	fmt.Println(reflectType(b))
	fmt.Println(reflectType(c))
	fmt.Println(reflectType(c2))

}

func CheckType(input interface{}) string {
	switch input.(type) {
	case string:
		return "string"
	case int:
		return "int"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	case chan struct{}:
		return "chan struct"
	}
	return "Cannot define type"
}

func reflectType(input interface{}) reflect.Type {
	return reflect.TypeOf(input)
}
