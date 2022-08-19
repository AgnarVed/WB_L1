package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	Name string
	Age  int
	Action
	Wait
}

type Wait struct {
	Time int
}

type Action struct {
	Walk int
}

func (a *Action) DoSomething() string {
	return "do exercises"
}

func main() {
	p1 := Human{
		Name: "Bob",
		Age:  23,
		Action: Action{
			Walk: 3,
		},
	}

	fmt.Printf("%s had walked %d kilometers this day\n", p1.Name, p1.Action.Walk)
	fmt.Printf("%s had walked %d kilometers this day and %s\n", p1.Name, p1.Walk, p1.Action.DoSomething())
	fmt.Printf("%d", p1.days(2))
}

func (a *Action) days(input int) int {
	result := input * 60
	return result
}

func (w *Wait) days(input int) int {
	result := input * 24
	return result
}

func (h *Human) days(input int) int {
	return h.Action.days(input)
}
