package main

import (
	"fmt"
	"strings"
)

// reverseWords принимает строку и возвращает её с перевёрнутыми словами
func reverseWords(input string) string {
	// Разделяем строку на слова
	words := strings.Fields(input)
	// Переворачиваем порядок слов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	// Объединяем слова обратно в строку
	return strings.Join(words, " ")
}

func main() {
	input := "snow dog sun"
	reversed := reverseWords(input)
	fmt.Printf("Исходная строка: %s\n", input)
	fmt.Printf("Строка с перевёрнутыми словами: %s\n", reversed)
}
