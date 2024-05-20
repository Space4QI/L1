package main

import (
	"fmt"
)

func checkType(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Переменная типа int")
	case string:
		fmt.Println("Переменная типа string")
	case bool:
		fmt.Println("Переменная типа bool")
	case chan int:
		fmt.Println("Переменная типа channel")
	default:
		fmt.Println("Не удалось определить тип переменной")
	}
}

func main() {
	var intValue int = 10
	var stringValue string = "hello"
	var boolValue bool = true
	var channelValue chan int

	fmt.Println("Проверка переменных:")
	checkType(intValue)
	checkType(stringValue)
	checkType(boolValue)
	checkType(channelValue)
}
