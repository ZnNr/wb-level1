package main

import "fmt"

/*
Разработать конвейер чисел.
Даны два канала:
в первый пишутся числа x из массива,
во второй – результат операции x*2.
После этого данные из второго канала должны выводиться в stdout.
То есть, организуйте конвейер из двух этапов с горутинами:
генерация чисел и их обработка.
Убедитесь, что чтение из второго канала корректно завершается.
*/

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	inChanel := make(chan int, len(numbers))
	outChanel := make(chan int, len(numbers))

	go func() { //для чтения чисел из массива и отправки в входящий канал
		for _, num := range numbers {
			inChanel <- num
		}
		close(inChanel)
	}()

	go func() { //для обработки чисел и отправки результатов исходящий канал
		for num := range inChanel {
			result := num * 2
			outChanel <- result
		}
		close(outChanel)
	}()

	for result := range outChanel {
		fmt.Println(result)
	}
}
