package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action
от родительской структуры Human (аналог наследования).
*/

// Human
type Human struct {
	Name string
	Age  int
}

// Методы Human
func (h *Human) Hello() {
	fmt.Printf("драствуйте, меня зовут %s, мне %d лет.\n", h.Name, h.Age)
}

func (h *Human) Running() {
	fmt.Printf("%s, бегит.\n", h.Name)
}

func (h *Human) Abdominal() {
	fmt.Printf("%s, пресс качат.\n", h.Name)
}

func (h *Human) Horizontal_bar() {
	fmt.Printf("%s, турник.\n", h.Name)
}

func (h *Human) Push_up() {
	fmt.Printf("%s, анжуманя.\n", h.Name)
}

// Action с встроенной структурой Human
type Action struct {
	Human // встраиваем Human
	Job   string
}

// Метод только для Action
func (a *Action) Work() {
	fmt.Printf("%s работает как %s.\n", a.Name, a.Job)
}

func main() {
	// экземпляр Action
	person := Action{
		Human: Human{Name: "Зиннур", Age: 40},
		Job:   "курэр",
	}
	person.Work() // вызов метода от Action

	//  методы Human через Action
	person.Hello()          // вызов унаследованного метода
	person.Abdominal()      // вызов унаследованного метода
	person.Running()        // вызов унаследованного метода
	person.Horizontal_bar() // вызов унаследованного метода
	person.Push_up()        // вызов унаследованного метода

	// бращение к полям Human напрямую
	fmt.Println("Имя:", person.Name)
}
