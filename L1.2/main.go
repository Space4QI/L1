package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	arr := []int{2, 4, 6, 8, 10}

	wg.Add(1)
	go calcSquare(arr, &wg)

	wg.Wait()
}

func calcSquare(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, num := range arr {
		square := num * num
		fmt.Println(square)
	}
}
