package main

import (
	"fmt"
)

func reverseString(input string) string {
	// Преобразуем строку в срез рун
	runes := []rune(input)
	// Переворачиваем срез рун
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	// Преобразуем срез рун обратно в строку
	return string(runes)
}

func main() {
	input := "главрыба"
	reversed := reverseString(input)
	fmt.Printf("Исходная строка: %s\n", input)
	fmt.Printf("Перевёрнутая строка: %s\n", reversed)
}
