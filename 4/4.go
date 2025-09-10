package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).
Выберите и обоснуйте способ завершения работы всех горутин-воркеров при получении сигнала прерывания.

Подсказка: можно использовать контекст (context.Context) или канал для оповещения о завершении.
*/

/*
наиболее простой, понятный способ завершения - Канал для сигналов + close() канала
за основу взят код предыдущей задачи
*/

func worker(id int, wg *sync.WaitGroup, jobs <-chan int) {
	defer wg.Done()

	// работает пока канал jobs не закрыт
	for n := range jobs {
		fmt.Printf("Worker %d processing: %d\n", id, n)
		time.Sleep(500 * time.Millisecond) // имитация работы
	}

	fmt.Printf("Worker %d: graceful shutdown completed\n", id)
}

func main() {
	var numWorkers int
	fmt.Print("Enter number of workers: ")

	if _, err := fmt.Scan(&numWorkers); err != nil || numWorkers < 1 {
		fmt.Println("Invalid number of workers")
		return
	}

	fmt.Printf("Starting %d workers...\n", numWorkers)
	fmt.Println("Press Ctrl+C to stop")

	// создается канал для jobs и сигналов
	jobs := make(chan int, 10)
	var wg sync.WaitGroup

	// канал для перехвата сигналов ОС
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// запуск workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, jobs)
	}

	// главная горутина пишет данные
	counter := 0
	for {
		select {
		case <-sigs:
			//  SIGINT - начинаем graceful shutdown
			fmt.Println("Received SIGINT. Starting graceful shutdown...")

			// Закрываем канал jobs - воркеры завершатся после обработки оставшихся задач
			close(jobs)

			// Ждем завершения всех воркеров
			wg.Wait()

			// Завершаем программу
			fmt.Println("All workers finished gracefully. Program exiting.")
			return

		default:
			// Продолжение работы
			jobs <- counter
			fmt.Printf("Main sent: %d\n", counter)
			counter++
			time.Sleep(200 * time.Millisecond)
		}
	}
}
