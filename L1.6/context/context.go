package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Воркер начал работу")
	select {
	case <-time.After(time.Second):
		fmt.Println("Воркер завершил работу")
	case <-ctx.Done():
		fmt.Println("Воркер был остановлен")
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx, &wg)
	time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()
	fmt.Println("Главная горутина завершила работу")
}
