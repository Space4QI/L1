package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func main() {
	fmt.Println("Аргументы командной строки:", os.Args)

	// Проверка наличия аргумента командной строки
	if len(os.Args) < 2 {
		fmt.Println("Использование: go run main.go <number_of_workers>")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Неправильный номер воркера:", err)
		return
	}

	fmt.Println("Число воркеров:", numWorkers)

	dataCh := make(chan int)
	var wg sync.WaitGroup

	// Создание воркеров
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataCh, &wg)
	}

	// Использование канала stopCh для сигнала остановки позволяет четко контролировать момент завершения работы горутины.
	stopCh := make(chan struct{})

	// Горутина для записи данных в канал
	go func() {
		i := 0
		for {
			select {
			case <-stopCh:
				close(dataCh)
				return
			case dataCh <- i:
				fmt.Printf("Запись данных в канал: %d\n", i)
				i++
			}
		}
	}()

	signCh := make(chan os.Signal, 1)

	// Данный сигнал завершения гарантирует, что все воркеры завершат свою работу и программа корректно завершится по сигналу от пользователя или системы
	signal.Notify(signCh, syscall.SIGINT, syscall.SIGTERM)

	<-signCh

	fmt.Println("Завершение работы")

	close(stopCh)
	wg.Wait()
	fmt.Println("Все воркеры завершили работу.")
}

func worker(id int, dataCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range dataCh {
		fmt.Printf("Воркер %d получил данные %d\n", id, data)
	}
	fmt.Printf("Воркер %d завершил работу\n", id)
}
