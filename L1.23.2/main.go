package main

import "fmt"

// Удаление элемента без сохранения порядка их следования
func removeElementWithoutOrder(s []int, index int) []int {
	// Заменяем удаляемый элемент последним элементом в срезе
	s[index] = s[len(s)-1]
	// Уменьшаем длину среза на 1
	return s[:len(s)-1]
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	index := 2 // Индекс элемента, который нужно удалить

	// Удаление элемента из среза
	result := removeElementWithoutOrder(slice, index)
	fmt.Println("Срез после удаления без сохранения порядка:", result)
}
