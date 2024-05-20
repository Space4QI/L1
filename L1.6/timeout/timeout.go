package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	fmt.Println("Работник начал работу")
	select {
	case <-time.After(time.Second):
		fmt.Println("Работник завершил работу")
	case <-ctx.Done():
		fmt.Println("Работник был остановлен")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	go worker(ctx)

	// Главная горутина завершится по тайм-ауту контекста
	<-ctx.Done()

	fmt.Println("Главная горутина завершила выполнение")
}
