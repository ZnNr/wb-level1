package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func intersect(a, b []int) []int {

	set := make(map[int]bool)
	for _, v := range a {
		set[v] = true
	}
	result := make(map[int]bool)
	var intersection []int
	for _, v := range b {
		if set[v] && !result[v] {
			result[v] = true
			intersection = append(intersection, v)
		}
	}
	return intersection
}
func printSet(arr []int, name string) {
	fmt.Print(name, " = {")
	for i, val := range arr {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Print("}")
}
func main() {
	//A := []int{1, 2, 3}
	//B := []int{2, 3, 4}

	scanner := bufio.NewScanner(os.Stdin)

	// просим длинну массива
	fmt.Print("Введите длину массива A: ")
	if !scanner.Scan() {
		fmt.Println("Ошибка ввода.")
		return
	}
	n, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil || n < 0 {
		fmt.Println("Некорректная длина.")
		return
	}
	
	fmt.Print("Введите длину массива B: ")
	if !scanner.Scan() {
		fmt.Println("Ошибка ввода.")
		return
	}
	m, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil || m < 0 {
		fmt.Println("Некорректная длина.")
		return
	}

	fmt.Printf("\nГенерируем массивы...\n")
	// Генерация случайных массивов
	A := make([]int, n)
	B := make([]int, m)
	// сслучайное заполнение числами от 1 до 10
	for i := range A {
		A[i] = rand.Intn(10) + 1
	}
	for i := range B {
		B[i] = rand.Intn(10) + 1
	}

	// вывод генераций на экран
	printSet(A, "A")
	fmt.Println()

	printSet(B, "B")
	fmt.Println()
	result := intersect(A, B)
	// вывод пересечений
	fmt.Print("Пересечение = {")
	for i, val := range result {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("}")
}
