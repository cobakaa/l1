//Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action
//от родительской структуры Human (аналог наследования).

package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) Walk() {
	fmt.Println("Human: Walk()")
}

func (h Human) Work() {
	fmt.Println("Human: Work()")
}

type Cat struct {
	name string
}

func (c Cat) Walk() {
	fmt.Println("Cat: Walk()")
}

// Встраивание методов

type Action struct {
	Human
	Cat
}

// Разпегение коллизии названий методов

func (a Action) Walk() {
	a.Cat.Walk()
}

func main() {
	h := Human{"Kirill", 16}
	c := Cat{"Anya"}
	a := Action{h, c}

	fmt.Println("---Human---")
	// Вызов методов структуры Human
	h.Walk()
	h.Work()

	fmt.Println()
	fmt.Println("---Action---")
	// Вызов методов стуктуры Action
	a.Walk()
	a.Work()
	a.Human.Walk()
}
