package main

import (
	"fmt"
)

// устанавливает i-й бит числа n в значение 1 или 0
func setBit(n int64, i uint, bit bool) int64 {
	if bit {
		// Устанавливаем i-й бит в 1
		n |= 1 << i
	} else {
		// Устанавливаем i-й бит в 0
		n &= ^(1 << i)
	}
	return n
}

func main() {
	var num int64 = 0 // начальное значение переменной
	i := uint(3)      // бит, который нужно изменить (начиная с 0)

	fmt.Printf("Начальное значение: %064b\n", num)

	// Устанавливаем i-й бит в 1
	num = setBit(num, i, true)
	fmt.Printf("Меняем бит номер %d на 1: %064b\n", i, num)

	// Устанавливаем i-й бит в 0
	num = setBit(num, i, false)
	fmt.Printf("Меняем бит номер %d на 0: %064b\n", i, num)
}
