package main

import "fmt"

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создание множества
	stringSet := make(map[string]struct{})

	// Добавляем строки в множество
	for _, str := range strings {
		stringSet[str] = struct{}{}
	}

	fmt.Println("Собственное множество:", stringSet)
}
