package main

import "fmt"

var justString string

func createHugeString(size int) string {
	// Пример создания строки заданного размера
	runes := make([]rune, size)
	for i := 0; i < size; i++ {
		runes[i] = 'a'
	}
	return string(runes)
}

func someFunc() {
	v := createHugeString(1 << 10)

	if len(v) < 100 {
		justString = v // если длина строки меньше 100, просто присваиваем всю строку
	} else {
		justString = string(v[:100]) // создаем новый срез и копируем его
	}
}

func main() {
	someFunc()
	fmt.Println(justString)
}
