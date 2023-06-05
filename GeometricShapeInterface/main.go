package main

import (
	"fmt"
	"math"
)

// Интерфейс Shape с методом Area()
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Структура Circle, реализующая интерфейс Shape
type Circle struct {
	Radius float64
}

// Метод Area для Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Метод Perimeter для Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Структура Rectangle, реализующая интерфейс Shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Метод Area для Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Метод Perimeter для Rectangle
func (r Rectangle) Perimeter() float64 {
	return r.Width*2 + r.Height*2
}

// Функция, принимающая интерфейс Shape
func PrintArea(s Shape) {
	fmt.Println("Площадь:", s.Area())
}

// Функция, принимающая интерфейс Shape
func PrintPerimeter(s Shape) {
	fmt.Println("Периметр:", s.Perimeter())
}

func main() {
	c := Circle{Radius: 5.0}
	r := Rectangle{Width: 4.0, Height: 3.0}

	PrintArea(c) // Выводит площадь круга
	PrintArea(r) // Выводит площадь прямоугольника

	PrintPerimeter(c) // Выводит периметр круга
	PrintPerimeter(r) // Выводит периметр прямоугольника
}
