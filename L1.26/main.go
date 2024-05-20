package main

import "fmt"

// isUnique возвращает true, если все символы в строке уникальные, и false в противном случае.
// Функция регистронезависима.
func isUnique(str string) bool {
	// Создаем map для отслеживания уникальных символов
	seen := make(map[rune]bool)

	// Проходим по всем символам в строке
	for _, char := range str {
		// Приводим символ к нижнему регистру
		char = toLower(char)
		// Проверяем, есть ли символ уже в map
		if seen[char] {
			return false
		}
		// Если символа нет в map, добавляем его
		seen[char] = true
	}
	return true
}

// toLower возвращает символ в нижнем регистре
func toLower(char rune) rune {
	if char >= 'A' && char <= 'Z' {
		return char + 32 // Разница между верхним и нижним регистром в таблице ASCII
	}
	return char
}

func main() {
	fmt.Println(isUnique("abcd"))      // true
	fmt.Println(isUnique("abCdefAaf")) // false
	fmt.Println(isUnique("aabcd"))     // false
}
