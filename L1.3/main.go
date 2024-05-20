package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	ch := make(chan int, len(arr))

	for _, num := range arr {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			square := num * num
			ch <- square
			fmt.Printf("%d*%d = %d\n", num, num, square)
		}(num)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	sum := 0
	for res := range ch {
		sum += res
	}

	fmt.Printf("Сумма квадратов = %d\n", sum)
}
