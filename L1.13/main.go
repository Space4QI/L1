package main

import "fmt"

func main() {
	a, b := 5, 10

	fmt.Println("Исходные значения:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// Обмен
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println("\nПосле обмена:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
