package main

import (
	"fmt"
	"math/big"
)

// Разработать программу,
// которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20

func main() {

	a := new(big.Int)
	a.SetString("2400000000000000000000000000000000000", 10)
	b := a
	c := big.NewInt(100000000000000)
	fmt.Println("Add: ", Add(a, c))
	fmt.Println("Mul: ", Mul(a, c))
	fmt.Println("Sub: ", Sub(a, c))
	fmt.Println("Div: ", Div(a, b))
}

func Add(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}
func Sub(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}
func Mul(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}
func Div(a *big.Int, b *big.Int) *big.Float {
	f1 := new(big.Float).SetInt(a)
	f2 := new(big.Float).SetInt(b)
	return new(big.Float).Quo(f1, f2)
}
