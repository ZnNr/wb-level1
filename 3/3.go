package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func worker(id int, wg *sync.WaitGroup, jobs <-chan int) {
	defer wg.Done()
	for n := range jobs {
		fmt.Printf("Worker %d: processing %d\n", id, n)
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Printf("Worker %d: shutdown complete\n", id)
}

func main() {
	fmt.Print("Enter number of workers: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	if len(input) > 0 && input[len(input)-1] == '\n' {
		input = input[:len(input)-1]
	}
	if len(input) > 0 && input[len(input)-1] == '\r' {
		input = input[:len(input)-1]
	}

	numWorkers, err := strconv.Atoi(input)
	if err != nil || numWorkers < 1 {
		fmt.Println("Error: please enter a valid positive integer")
		return
	}

	fmt.Printf("Starting %d workers...\n", numWorkers)

	jobs := make(chan int, 15)
	var wg sync.WaitGroup
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, jobs)
	}

	counter := 0
	for {
		select {
		case <-sigs:
			fmt.Println("\nShutting down...")
			close(jobs)
			wg.Wait()
			fmt.Println("All workers completed!")
			return
		default:
			jobs <- counter
			fmt.Printf("Sent: %d\n", counter)
			counter++
			time.Sleep(150 * time.Millisecond)
		}
	}
}
