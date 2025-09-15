package main

import "fmt"

/*
Дана переменная типа int64. Разработать программу, которая устанавливает i-й бит этого числа в 1 или 0.

Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).

Подсказка: используйте битовые операции (|, &^).
*/

func setBit(num int64, pos uint, bit int) int64 {
	if bit == 1 {
		return num | (1 << pos) // Установка бита в 1
	}
	return num & ^(1 << pos) // Установка бита в 0
}

func main() {
	var num int64
	var pos uint
	var bit int // ввод 0 или 1 корректно работал
	// Ввод натурального числа
	fmt.Print("Enter natural number: ")
	_, err := fmt.Scan(&num)
	if err != nil || num < 0 {
		fmt.Println("Please enter a positive natural number")
		return
	}

	// Ввод позиции бита
	fmt.Print("Enter bit position (starting from 0): ")
	_, err = fmt.Scan(&pos)
	if err != nil {
		fmt.Println("Invalid position:", err)
		return
	}
	if pos > 63 { // int64
		fmt.Println("Position must be between 0 and 63")
		return
	}

	// Ввод бита (0 или 1)
	fmt.Print("Enter 0 or 1: ")
	_, err = fmt.Scan(&bit)
	if err != nil || (bit != 0 && bit != 1) {
		fmt.Println("Please enter 0 or 1!")
		return
	}

	result := setBit(num, pos, bit)
	fmt.Printf("%d (%b) -> установка %d-го бита в %b -> %d (%b)\n",
		num, num, pos, bit, result, result)
}
