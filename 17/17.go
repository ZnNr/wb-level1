package main

import (
	"fmt"
)

/*
Реализовать алгоритм бинарного поиска встроенными методами языка. Функция должна принимать отсортированный слайс и искомый элемент, возвращать индекс элемента или -1, если элемент не найден.

Подсказка: можно реализовать рекурсивно или итеративно, используя цикл for.
*/

// binarySearch получает отсортированный срез и ищет индекс элемента в отсортированном слайсе
func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid // элемент найден
		} else if arr[mid] < target {
			left = mid + 1 // ищем в правой половине
		} else {
			right = mid - 1 // ищем в левой половине
		}
	}

	return -1 // элемент не найден
}

func main() {
	arr := []int{-5, -1, 0, 3, 7, 9, 12, 15, 20}
	target := 6

	index := binarySearch(arr, target)
	if index != -1 {
		fmt.Printf("Элемент %d найден на индексе %d.\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден в массиве.\n", target)
	}
}
