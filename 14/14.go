package main

import "fmt"

/*
Разработать программу, которая в runtime способна определить тип переменной, переданной в неё (на вход подаётся interface{}). Типы, которые нужно распознавать: int, string, bool, chan (канал).

Подсказка: оператор типа switch v.(type) поможет в решении.
*/

func detectType(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Printf("Тип: int, значение: %d\n", val)
	case string:
		fmt.Printf("Тип: string, значение: %s\n", val)
	case bool:
		fmt.Printf("Тип: bool, значение: %t\n", val)
	case chan int:
		fmt.Printf("Тип: chan int, адрес канала: %p\n", val)
	case chan string:
		fmt.Printf("Тип: chan string, адрес канала: %p\n", val)
	case chan bool:
		fmt.Printf("Тип: chan bool, адрес канала: %p\n", val)
	case chan interface{}:
		fmt.Printf("Тип: chan interface{}, адрес канала: %p\n", val)
	default:
		fmt.Printf("Неизвестный тип: %T\n", v)
	}
}

func main() {
	// Примеры значений разных типов
	var chInt = make(chan int)
	var chStr = make(chan string)
	var chBool = make(chan bool)

	detectType(123456)                 // int
	detectType("qwerty")               // string
	detectType(true)                   // bool
	detectType(chInt)                  // chan int
	detectType(chStr)                  // chan string
	detectType(chBool)                 // chan bool
	detectType(make(chan interface{})) // chan interface{}
	detectType(3.14)                   // float64 — не поддерживается, значит default
}
