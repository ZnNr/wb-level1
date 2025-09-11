package main

import (
	"context"
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.

Классические подходы: выход по условию, через канал уведомления, через контекст, прекращение работы runtime.Goexit() и др.

Продемонстрируйте каждый способ в отдельном фрагменте кода.
*/
func main() {
	fmt.Println("Демонстрация способов остановки горутины\n")

	example1_ChannelSignal()       // через сигнал в канал
	example2_ContextCancel()       // через контекст с отменой
	example3_AtomFlag()            // флаг + sync/atomic
	example4_CloseChannel()        // через закрытие канала
	example5_RuntimeGoexit()       // через runtime.Goexit()
	example6_ContextWithTimeout()  // через контекст с таймаутом
	example7_BufferedChannelFlag() // через буферизированный канал флаг
	example8_ContextWithDeadline() // через контекст с дедлайном
	example9_SimpleFlag()          // простой флаг (без atomic)

	example_NaturalReturn() // естественное завершение

	fmt.Println("\nВсе примеры завершены.")
}

// Пример 1: Остановка через канал сигнала
func example1_ChannelSignal() {
	fmt.Println("[пример 1] Остановка через канал (канал-сигнал)")
	stop := make(chan struct{})

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Горутина остановлена через канал")
				return
			default:
				fmt.Println("Работаю с каналом...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	close(stop)
	time.Sleep(100 * time.Millisecond)
}

// Пример 2: Остановка через context
func example2_ContextCancel() {
	fmt.Println("\n[пример 2] Остановка через context.Context")
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("Горутина остановлена: %v\n", ctx.Err())
				return
			default:
				fmt.Println("Что-то происходит с контекстом...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(100 * time.Millisecond)
}

// Пример 3: Остановка через атомарный флаг
func example3_AtomFlag() {
	fmt.Println("\n[пример 3] Остановка через атомарный флаг")
	var running atomic.Bool
	running.Store(true)

	go func() {
		for running.Load() {
			fmt.Println("Работа работа...  флагом...")
			time.Sleep(1 * time.Second)
		}
		fmt.Println("Горутина с флагом остановлена")
	}()

	time.Sleep(3 * time.Second)
	running.Store(false)
	time.Sleep(100 * time.Millisecond)
}

// Пример 4: Остановка через закрытие канала (range)
func example4_CloseChannel() {
	fmt.Println("\n[пример 4] Остановка при закрытии канала (range)")
	data := make(chan string)

	go func() {
		for msg := range data { // автоматически завершится при close(data)
			fmt.Printf("Получено: %s\n", msg)
		}
		fmt.Println("Горутина завершилась после закрытия канала")
	}()

	data <- "привет"
	data <- "мирок"
	close(data)
	time.Sleep(100 * time.Millisecond)
}

// Пример 5: Остановка через runtime.Goexit()
func example5_RuntimeGoexit() {
	fmt.Println("\n[пример 5] Остановка через runtime.Goexit()")

	go func() {
		defer func() {
			fmt.Println("  defer выполнился (Goexit не отменяет defer)")
		}()

		fmt.Println("Начали работу...")
		time.Sleep(1 * time.Second)

		fmt.Println("Вызыв runtime.Goexit()")
		runtime.Goexit() // немедленно завершает эту горутину

		fmt.Println("Это не напечатается") // unreachable
	}()

	time.Sleep(3 * time.Second)
}

// Пример 6: context.WithTimeout — автоматический таймаут
func example6_ContextWithTimeout() {
	fmt.Println("\n[6] Остановка через context.WithTimeout (3 сек)")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // хорошая практика

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("Контекст завершён: %v (таймаут)\n", ctx.Err())
				return
			default:
				fmt.Println("Работа работа... (WithTimeout)")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(4 * time.Second)
}

// Пример 7: Буферизированный канал-флаг
func example7_BufferedChannelFlag() {
	fmt.Println("\n[пример 7] Остановка через буферизированный канал")
	stop := make(chan bool, 1) // буфер на 1, чтобы не блокировать

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Получен сигнал остановки (буферизированный канал)")
				return
			default:
				fmt.Println("Работа работа... (буферизированный канал)")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	stop <- true // отправляем сигнал
	time.Sleep(100 * time.Millisecond)
}

// Пример 8: context.WithDeadline — остановка в конкретное время
func example8_ContextWithDeadline() {
	fmt.Println("\n[прмерчик 8] Остановка через context.WithDeadline (через 3 секунды)")

	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("Контекст завершён: %v (дошло до дедлайна)\n", ctx.Err())
				return
			default:
				fmt.Println(" еще работа... (WithDeadline)")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(4 * time.Second)
}

// --- Пример 9: Простой флаг (без atomic) ---
func example9_SimpleFlag() {
	fmt.Println("\n[пример 9] Остановка через простой флаг (без atomic)")
	running := true

	go func() {
		for running {
			fmt.Println("Работаю с простым флагом...")
			time.Sleep(1 * time.Second)
		}
		fmt.Println("Горутина с простым флагом остановлена")
	}()

	time.Sleep(3 * time.Second)
	running = false
	time.Sleep(100 * time.Millisecond)
}

// --- Пример: Естественное завершение ---
// горутина останавливается выполнив весь свой код до конца (достигнув return или конца функции).
func example_NaturalReturn() {
	fmt.Println("\n[ессественное завершение] Естественное завершение (return)")

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("Естественно работаю?...естественно!")
			time.Sleep(1 * time.Second)
		}
		fmt.Println("Естественное завершение выполнено поработали и хватит")
	}()

	time.Sleep(4 * time.Second)
}
