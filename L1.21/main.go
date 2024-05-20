package main

import (
	"fmt"
)

// Point2D 2D координаты
type Point2D struct {
	X, Y int
}

// Point3D 3D координаты
type Point3D struct {
	X, Y, Z int
}

// Adapter Адаптер для преобразования 2D координат в 3D координаты
type Adapter struct {
	Point2D
}

// To3D Метод для преобразования 2D координат в 3D координаты
func (a *Adapter) To3D() Point3D {
	// Просто добавляем третью координату Z, принимая её равной нулю
	return Point3D{a.X, a.Y, 0}
}

func main() {
	// Создаем объект адаптера
	adapter := Adapter{Point2D{10, 20}}

	// Преобразуем 2D координаты в 3D координаты с помощью адаптера
	point3D := adapter.To3D()

	fmt.Println("Преобразованные 3D координаты:", point3D)
}
