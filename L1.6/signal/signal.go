package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker() {
	fmt.Println("Работник начал работу")
	time.Sleep(time.Second)
	fmt.Println("Работник завершил работу")
}

func main() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigCh
		fmt.Println("Получен сигнал для остановки")
		os.Exit(1)
	}()

	go worker()

	fmt.Println("Главная горутина запущена")

	// Ожидание сигнала завершения
	<-sigCh
}
