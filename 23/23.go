package main

import "fmt"

/*
Удалить i-ый элемент из слайса. Продемонстрируйте корректное удаление без утечки памяти.

Подсказка: можно сдвинуть хвост слайса на место удаляемого элемента (copy(slice[i:], slice[i+1:])) и уменьшить длину слайса на 1.
*/

func removeElement(slice []int, i int) []int {
	if i < 0 || i > len(slice)-1 {
		fmt.Println("Индекс вне диапазона.")
		return slice
	}
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	indexToRemove := 2
	removedValue := slice[indexToRemove]
	fmt.Println("Первоначальный слайс:", slice)

	fmt.Printf("Удаляем элемент с индексом %d: значение = %d\n", indexToRemove, removedValue)
	newSlice := removeElement(slice, indexToRemove)
	fmt.Println("Слайс после удаления:", newSlice)

}
