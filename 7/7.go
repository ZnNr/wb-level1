package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Реализовать безопасную для конкуренции запись данных в структуру map.

Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).

Проверьте работу кода на гонки (util go run -race).
*/

// SafeMap - потокобезопасная map с RWMutex
type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{data: make(map[string]int)}
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return sm.data[key], true
}

func (sm *SafeMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key)
}

func (sm *SafeMap) Len() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return len(sm.data)
}

// Тест конкурентной записи и чтения
func testConcurrentAccess() {
	fmt.Println("=== Тест конкурентного доступа ===")
	sm := NewSafeMap()
	var wg sync.WaitGroup

	// Запись из 100 горутин
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			sm.Set(fmt.Sprintf("key%d", idx), idx)
		}(i)
	}

	// Чтение из 50 горутин
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			sm.Get(fmt.Sprintf("key%d", idx%50))
		}(i)
	}

	wg.Wait()
	fmt.Printf("Записано элементов: %d\n", sm.Len())
}

// Тест на гонки данных
func testDataRaces() {
	fmt.Println("\n=== Тест на гонки данных ===")
	sm := NewSafeMap()
	var wg sync.WaitGroup

	// Интенсивная конкурентная работа
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			key := fmt.Sprintf("race_key%d", idx%100)
			sm.Set(key, idx)
			sm.Get(key)
			if idx%10 == 0 {
				sm.Delete(key)
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("Финальный размер map: %d\n", sm.Len())
}

// Сравнение с обычной map (НЕбезопасно!)
func demonstrateUnsafeMap() {
	fmt.Println("\n=== Демонстрация UNSAFE map ===")
	unsafeMap := make(map[string]int)
	var wg sync.WaitGroup

	// Эта часть вызовет гонки данных!
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			key := fmt.Sprintf("unsafe%d", idx)
			unsafeMap[key] = idx // ГОНКА ДАННЫХ!
		}(i)
	}

	wg.Wait()
	fmt.Printf("Unsafe map size: %d (данные могут быть повреждены)\n", len(unsafeMap))
}

// Тест sync.Map (встроенная потокобезопасная map)
func testSyncMap() {
	fmt.Println("\n=== Тест sync.Map ===")
	var sm sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			key := fmt.Sprintf("sync_key%d", idx)
			sm.Store(key, idx)

			// Чтение
			if val, ok := sm.Load(key); ok {
				_ = val // Используем значение
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Sync.Map тест завершен")
}

// Проверка целостности данных
func testDataIntegrity() {
	fmt.Println("\n=== Проверка целостности данных ===")
	sm := NewSafeMap()
	var wg sync.WaitGroup

	// Запись
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			sm.Set(fmt.Sprintf("test%d", idx), idx*2)
		}(i)
	}
	wg.Wait()

	// Проверка
	errors := 0
	for i := 0; i < 200; i++ {
		if val, ok := sm.Get(fmt.Sprintf("test%d", i)); !ok || val != i*2 {
			errors++
		}
	}

	fmt.Printf("Ошибок целостности: %d/200\n", errors)
}

func main() {
	fmt.Println("go run -race выполнять из unix или wsl")
	fmt.Println("Для проверки на гонки: go run -race 7/7.go")

	start := time.Now()

	testConcurrentAccess()
	testDataRaces()
	testDataIntegrity()
	testSyncMap()
	demonstrateUnsafeMap()

	fmt.Printf("\nВсе тесты завершены за: %v\n", time.Since(start))
	fmt.Println("ВЫВОДЫ:")
	fmt.Println("SafeMap безопасна для конкурентного использования")
	fmt.Println("Sync.Map - готовая потокобезопасная реализация")
	fmt.Println("Обычная map вызывает гонки данных")
}
