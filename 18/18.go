package main

import (
	"fmt"
	"sync"
)

/*

Реализовать структуру-счётчик, которая будет инкрементироваться в конкурентной среде (т.е. из нескольких горутин).
По завершению программы структура должна выводить итоговое значение счётчика.

Подсказка: вам понадобится механизм синхронизации, например, sync.Mutex или sync/Atomic для безопасного инкремента.
*/

// counter
type Counter struct {
	mu    sync.Mutex
	value int
}

// увеличивалка счетчика
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// возвращатель значения счетчика
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}
	for j := 0; j < 10; j++ {
		wg.Go(func() { //Метод WaitGroup.Go v1.25
			for i := 0; i < 1000; i++ {
				counter.Inc()
			}
		})
	}
	wg.Wait()
	fmt.Println("значение счетчика: ", counter.Value())
}
