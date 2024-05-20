package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	count int64
}

func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.count, 1)
}

func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.count)
}

func main() {
	var wg sync.WaitGroup
	counter := &AtomicCounter{}

	// Запускаем несколько горутин для инкрементации счетчика
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Println("Final counter value:", counter.Value())
}
