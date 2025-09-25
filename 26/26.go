package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке встречаются один раз
(т.е. строка состоит из уникальных символов).

Вывод: true, если все символы уникальны, false, если есть повторения.
Проверка должна быть регистронезависимой, т.е. символы в разных регистрах считать одинаковыми.

Например: "abcd" -> true, "abCdefAaf" -> false (повторяются a/A), "aabcd" -> false.

Подумайте, какой структурой данных удобно воспользоваться для проверки условия.
*/
func uniqSymbols(s string) bool {
	charSet := make(map[rune]struct{})
	s = strings.ToLower(s)
	for _, char := range s {
		if _, exists := charSet[char]; exists {
			return false
		}
		charSet[char] = struct{}{}
	}
	return true
}
func main() {
	testStrings := []string{"abcd", "abCdefAaf", "aabcd", "AaCcDd"}

	for _, str := range testStrings {
		result := uniqSymbols(str)
		fmt.Printf("Строка: %s, Уникальные символы: %t\n", str, result)
	}
}
