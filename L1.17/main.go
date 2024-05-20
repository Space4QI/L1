package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 4

	// Использование sort.Search для бинарного поиска
	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	// Проверка, найден ли элемент
	if index < len(arr) && arr[index] == target {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found\n", target)
	}
}
