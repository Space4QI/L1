package main

import (
	"fmt"
)

// функция для записи чисел из массива в первый канал
func writer(numbers []int, ch chan int) {
	for _, n := range numbers {
		ch <- n
	}
	close(ch)
}

// функция для чтения из первого канала, умножения на 2 и записи во второй канал
func multiplier(in chan int, out chan int) {
	for n := range in {
		out <- n * 2
	}
	close(out)
}

// функция для чтения из второго канала и вывода в stdout
func printer(ch chan int) {
	for n := range ch {
		fmt.Println(n)
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go writer(numbers, ch1)
	go multiplier(ch1, ch2)
	printer(ch2)
}
