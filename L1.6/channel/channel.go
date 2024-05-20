package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Работник начал работу")
	time.Sleep(time.Second)
	fmt.Println("Работник закончил работу")
	done <- true
}

func main() {
	done := make(chan bool)
	go worker(done)
	<-done
	fmt.Println("Завершение работы")
}
