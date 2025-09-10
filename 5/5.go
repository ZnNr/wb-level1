package main

import (
	"fmt"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала – читать эти значения. По истечении N секунд программа должна завершаться.

Подсказка: используйте time.After или таймер для ограничения времени работы.
*/

func producer(ch chan<- int, done <-chan struct{}) {
	num := 0
	for {
		select {
		case ch <- num:
			fmt.Printf("Производитель отправил: %d\n", num)
			num++
			time.Sleep(1 * time.Second)
		case <-done:
			close(ch) // Закрываем канал при сигнале
			return
		}
	}
}

func consumer(ch <-chan int, done <-chan struct{}) {
	for {
		select {
		case item, ok := <-ch:
			if !ok {
				return // канал закрыт
			}
			fmt.Printf("Потребитель получил: %d\n", item)
		case <-done:
			return // время вышло
		}
	}
}

func main() {
	const N = 5 * time.Second

	ch := make(chan int)
	done := make(chan struct{}) // общий сигнал завершения

	// Отправляем сигнал через N секунд
	go func() {
		time.Sleep(N)
		close(done)
	}()

	go producer(ch, done)
	go consumer(ch, done)

	// ожидапние  чтобы потребитель мог обработать последнее значение
	time.Sleep(N + 1*time.Second)
	fmt.Println("Программа завершена.")
}
