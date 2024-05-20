package main

import (
	"fmt"
	"sync"
)

type myMap struct {
	mu    sync.Mutex
	items map[int]int
}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	m := myMap{
		items: make(map[int]int),
	}

	wg.Add(2)
	go writer(ch, &wg, &m)
	go writer(ch, &wg, &m)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for val := range ch {
		fmt.Printf("Received: %d\n", val)
	}

	fmt.Println("Map content:", m.items)
}

func writer(ch chan int, wg *sync.WaitGroup, m *myMap) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		m.mu.Lock()
		if _, ok := m.items[i]; !ok {
			m.items[i] = i
			ch <- i
		}
		m.mu.Unlock()
	}
}
