package main

import "fmt"

// Удаление элемента с сохранением порядка их следования
func removeElement(s []int, index int) []int {
	// Удаляем элемент по индексу
	return append(s[:index], s[index+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	index := 2 // Индекс элемента, который нужно удалить

	// Удаление элемента из среза
	result := removeElement(slice, index)
	fmt.Println("Срез после удаления:", result)
}
