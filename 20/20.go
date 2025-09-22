package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает порядок слов в строке.

Пример: входная строка:

«snow dog sun», выход: «sun dog snow».

Считайте, что слова разделяются одиночным пробелом.
Постарайтесь не использовать дополнительные срезы, а выполнять операцию «на месте».
*/

func reverseWords(s string) string {
	words := strings.Fields(s)
	n := len(words)
	for i := 0; i < n/2; i++ {
		words[i], words[n-1-i] = words[n-1-i], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	inpt := "snow dog sun"
	reversed := reverseWords(inpt)
	fmt.Printf("Origin: %s\n", inpt)
	fmt.Printf("Result: %s\n", reversed)
}
