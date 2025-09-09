package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Написать программу, которая конкурентно рассчитает
значение квадратов чисел взятых из массива (2,4,6,8,10)
и выведет их квадраты в stdout.
*/

func calcSquarePants(wg *sync.WaitGroup, num int) {
	defer wg.Done()
	squarePants := num * num
	fmt.Printf("Время: %s | Квадрат числа %d равен %d\n", time.Now().Format("15:04:05.999999999"), num, squarePants)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go calcSquarePants(&wg, num)
	}

	wg.Wait()
}
