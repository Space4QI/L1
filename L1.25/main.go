package main

import (
	"fmt"
	"time"
)

// Sleep задерживает выполнение программы на заданное количество секунд.
func sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println("Начало программы")
	sleep(3) // Задержка на 3 секунды
	fmt.Println("Прошло 3 секунды")
}
