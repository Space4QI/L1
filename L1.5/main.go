package main

import (
	"fmt"
	"time"
)

func sender(ch chan<- int, done chan<- struct{}) {
	defer close(ch)
	for i := 0; i < 5; i++ {
		fmt.Println("Отправка значения в канал:", i)
		ch <- i
		fmt.Println("Значение отправлено в канал:", i)
		time.Sleep(time.Second) // Ожидание одной секунды перед отправкой следующего значения
	}
	close(done) // Сигнализация о завершении отправки
}

func receiver(ch <-chan int, done <-chan struct{}) {
	for {
		select {
		case num, ok := <-ch:
			if !ok {
				fmt.Println("Горутина чтения завершила работу.")
				return
			}
			fmt.Println("Принято значение из канала:", num)
		case <-done:
			fmt.Println("Принял сигнал о завершении отправки.")
			return
		}
	}
}

func main() {
	ch := make(chan int)
	done := make(chan struct{})

	go sender(ch, done)
	go receiver(ch, done)

	// Ожидаем завершения отправки
	<-done

	fmt.Println("Программа завершена.")
}
