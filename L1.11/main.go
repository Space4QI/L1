package main

import (
	"fmt"
)

// intersect возвращает пересечение двух множеств
func intersect(set1, set2 []int) []int {
	// Создаем map для хранения элементов первого множества
	setMap := make(map[int]struct{})
	for _, v := range set1 {
		setMap[v] = struct{}{}
	}

	// Создаем срез для хранения пересечения
	var intersection []int
	for _, v := range set2 {
		if _, found := setMap[v]; found {
			intersection = append(intersection, v)
		}
	}

	return intersection
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{4, 5, 6, 7, 8}

	result := intersect(set1, set2)

	fmt.Println("Пересечение множеств:", result)
}
