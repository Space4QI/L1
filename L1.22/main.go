package main

import (
	"fmt"
	"math"
)

func main() {
	a := math.MaxInt32 + 1 // Значение a > 2^20
	b := math.MaxInt32 + 1 // Значение b > 2^20

	// Сложение
	sum := a + b
	fmt.Printf("Сумма a и b: %d\n", sum)

	// Вычитание
	subtraction := a - b
	fmt.Printf("Разность a и b: %d\n", subtraction)

	// Умножение
	multiplication := int64(a) * int64(b)
	fmt.Printf("Произведение a и b: %d\n", multiplication)

	// Деление
	division := float64(a) / float64(b)
	fmt.Printf("Частное a и b: %.2f\n", division)
}
