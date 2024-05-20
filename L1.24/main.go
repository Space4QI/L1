package main

import (
	"fmt"
	"math"
)

// Point Структура для представления точки в двумерном пространстве
type Point struct {
	x, y float64
}

// NewPoint Конструктор для создания новых точек
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// DistanceTo Метод для нахождения расстояния между текущей точкой и другой точкой
func (p Point) DistanceTo(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаем две точки
	point1 := NewPoint(1, 1)
	point2 := NewPoint(4, 5)

	// Находим расстояние между ними
	distance := point1.DistanceTo(point2)
	fmt.Printf("Расстояние между точкой 1 и точкой 2: %.2f\n", distance)
}
