package main

import (
	"fmt"
	"math"
	"strconv"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0

func main() {
	a := int64(145)
	fmt.Println(strconv.FormatInt(a, 2))
	fmt.Println(changeBit(3, a))

}
func changeBit(index float64, input int64) string {
	addVal := math.Pow(2, index)
	addVal2 := int64(addVal)
	val := input ^ addVal2
	result := strconv.FormatInt(val, 2)
	return result
}
