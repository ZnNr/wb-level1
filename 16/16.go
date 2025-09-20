package main

import "fmt"

/*
Реализовать алгоритм быстрой сортировки массива встроенными средствами языка. Можно использовать рекурсию.

Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел. Для выбора опорного элемента можно взять середину или первый элемент.
*/

//алгоритм быстрой сортировки (quicksort) с использованием встроенных методов.

// quickSort выполняет быструю сортировку массива с помощью рекурсии
func quickSort(array []int) []int {
	if len(array) < 2 {
		return array // массив из 0 или 1 элемента уже отсортирован
	}

	pivot := array[len(array)/2] // выбираем опорный/средний элемент
	left := []int{}
	right := []int{}
	var equal []int // массив для элементов, равных опорному

	for _, value := range array {
		if value < pivot {
			left = append(left, value) // элементы меньше опорного
		} else if value > pivot {
			right = append(right, value) // элементы больше опорного
		} else {
			equal = append(equal, value) // элементы равны опорному
		}
	}

	// рекурсивное применение быстрой сортировки и объединение массивов
	return append(append(quickSort(left), equal...), quickSort(right)...)
}

func main() {
	// Пример массива для сортировки
	unsorted := []int{64, 34, 25, 12, 22, 11, 90, 5}
	fmt.Println("Исходный массив:", unsorted)

	// Сортировка массива
	sortedArray := quickSort(unsorted)
	fmt.Println("Отсортированный массив:", sortedArray)
}
