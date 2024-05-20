package main

import (
	"fmt"
	"sync"
)

// Counter представляет собой структуру с мьютексом и счётчиком
type Counter struct {
	mu    sync.Mutex
	count int
}

// Increment увеличивает значение счётчика на единицу
func (c *Counter) Increment() {
	c.mu.Lock()         // Блокируем мьютекс
	defer c.mu.Unlock() // Разблокируем мьютекс при выходе из функции
	c.count++
}

// Value возвращает текущее значение счётчика
func (c *Counter) Value() int {
	c.mu.Lock()         // Блокируем мьютекс
	defer c.mu.Unlock() // Разблокируем мьютекс при выходе из функции
	return c.count
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	const numGoroutines = 100
	const numIncrements = 1000

	// Запускаем numGoroutines горутин
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done() // Уменьшаем счётчик WaitGroup при завершении горутины
			// Каждая горутина увеличивает счётчик numIncrements раз
			for j := 0; j < numIncrements; j++ {
				counter.Increment()
			}
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
	fmt.Println("Final Counter:", counter.Value())
}
