package main

import (
	"fmt"
	"time"
)

func worker() {
	fmt.Println("Работник начал работу")
	select {
	case <-time.After(time.Second):
		fmt.Println("Работник завершил работу")
	}
}

func main() {
	go worker()

	// Ждем определенное время, затем просто завершаем главную горутину
	time.Sleep(1500 * time.Millisecond)

	fmt.Println("Главная горутина завершила выполнение")
}
