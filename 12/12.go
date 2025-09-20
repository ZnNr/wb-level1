package main

import (
	"fmt"
	"sort"
)

/*
Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.

Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.
*/

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	//ключ — строка (string)
	//значение — тип struct{} — пустая структура, занимает 0 байт памяти
	set := make(map[string]struct{})

	// ключ в мапе всегда уникальный, поэтому удаляет дубликаты
	for _, str := range strings {
		set[str] = struct{}{} // Добавляем строку в множество
	}
	//красивый вывод в скобочках
	fmt.Print("Элементы множества: {")

	keys := make([]string, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i, k := range keys {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(k)
	}
	fmt.Println("}")
}
