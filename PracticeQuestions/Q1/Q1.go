package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

// Структура Человек
type Human struct {
	Name string
	Age  int
	Action
	Wait
}

// Ожидание
type Wait struct {
	Time int
}

// Действие
type Action struct {
	Walk int
}

// Сделать действие (упражнения)
func (a *Action) DoSomething() string {
	return "do exercises"
}

func main() {
	// Человек р1
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

// Например, нужно узнать, сколько времени нужно делать упражнения, чтобы достичь результата
func (a *Action) days(input int) int {
	result := input * 60
	return result
}

// Но также узнать, сколько нужно времени, чтобы успеть отдохнуть
func (w *Wait) days(input int) int {
	result := input * 24
	return result
}

// Указываем, что при запросе "days" возвращалось кол-во времени, в течении которого нужно выполнять упражнения,
// т.к. происходит наложение ( ambiguous reference )
func (h *Human) days(input int) int {
	return h.Action.days(input)
}
