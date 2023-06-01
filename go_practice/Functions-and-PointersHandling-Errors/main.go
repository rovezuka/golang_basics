package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	printCircleArea(-2) // калькулятор площади круга
	printCircleArea(2)

	// указатели

	x := 10
	p := &x
	fmt.Println("Значение x: ", x)
	fmt.Println("Адрес переменной x:", &x) // адрес переменной x
	fmt.Println("Значение *p", *p)         // распаковка адреса переменной, возвращается ее значение
	increment(&x)                          // изменения значения переменной по адресу
	fmt.Println("Значение переменной x после функции increment:", x)

	var n *int
	fmt.Println(n) // переменная без ссылки в памяти

	mathAndErrors() // библиотеки math и errors

	a := 5 // значения для определения площадей и периметров фигур
	b := 10
	c := 7
	h := 3
	rectangleArea := calculateRectangleArea(a, b)
	fmt.Println("Площадь прямоугольника:", rectangleArea)

	triangleArea := calculateTriangleArea(float32(a), float32(h))
	fmt.Println("Площадь треугольника:", triangleArea)

	rectanglePerimeter := calculateRectanlgePerimeter(a, b)
	fmt.Println("Периметр прямоугольника:", rectanglePerimeter)

	trianglePerimeter := calculateTrianglePerimeter(a, b, c)
	fmt.Println("Периметр треугольника:", trianglePerimeter)

	conditionalStatements() // условные конструкции

	signposts() // new() и указатели

}

func increment(p *int) {
	*p += 1
}

func printCircleArea(radius int) { // функция вывода площади круга

	circleArea, err := calculateCircleArea(radius)
	if err != nil { // проверка, что введено положительное значение
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Радиус круга: %d см.\n", radius)
	fmt.Println("Формула для расчета площади круга: A=pi*r^2")

	fmt.Printf("Площадь круга: %f см. кв.\n", circleArea)
}

func calculateCircleArea(radius int) (float32, error) { // функция подсчета площади круга
	if radius <= 0 {
		return float32(0), errors.New("Радиус круга не может быть отрицательным!")
	}
	return float32(radius) * float32(radius) * math.Pi, nil
}

func calculateTriangleArea(base float32, higher float32) float32 { // функция подсчета площади треугольника
	area := float32(1) / float32(2) * base * higher
	return area
}

func calculateRectangleArea(length int, width int) int { // функция подсчета площади прямоугольника
	area := length * width
	return area
}

func calculateTrianglePerimeter(a int, b int, c int) int { // функция подсчета периметра треугольника
	perimeter := a + b + c
	return perimeter
}

func calculateRectanlgePerimeter(length int, width int) int { // функция подсчета периметра треугольника
	perimeter := (length + width) * 2
	return perimeter
}

func mathAndErrors() { // пакеты math и errors
	err := errors.New("Это пример ошибки") // создание новой ошибки с помощью функции New
	fmt.Println(err)

	err1 := errors.New("Ошибка 1")
	err2 := errors.New("Ошибка 2")

	if err1.Error() == err2.Error() { // сравнение ошибок
		fmt.Println("Ошибки равны")
	} else {
		fmt.Println("Ошибки не равны")
	}

	x := 2.0
	y := 3.0

	fmt.Println("Квадратный корень из", x, ":", math.Sqrt(x)) // методы библиотеки math
	fmt.Println("Синус", y, ":", math.Sin(y))

	fmt.Println("Значение π:", math.Pi) // математические константы
	fmt.Println("Значение е:", math.E)

	x = 3.7
	y = 2.3

	fmt.Println("Округление", x, "вниз:", math.Floor(x)) // округление в библиотеке math
	fmt.Println("Округление", y, "вверх:", math.Ceil(y))

}

func conditionalStatements() { // функция, показывающая работу с условными конструкциями
	score := 85

	if score >= 90 {
		fmt.Println("Отличный результат!")
	} else if score >= 80 {
		fmt.Println("Хороший результат.")
	} else if score >= 70 {
		fmt.Println("Средний результат.")
	} else {
		fmt.Println("Низкий результат.")
	}
}

func signposts() { // функция для работы с new() и указателями
	/* Функция new() полезна, когда вам нужно выделить память для переменных или структур,
	особенно если они являются типами данных с указателем или если вы хотите инициализировать их
	значениями по умолчанию.
	*/
	// Создание указателя на целочисленную переменную
	numPtr := new(int)
	fmt.Println("Значение по умолчанию:", *numPtr) // Выводит 0

	// Создание указателя на строку
	strPtr := new(string)
	fmt.Println("Значение по умолчанию:", *strPtr) // Выводит пустую строку

	// Создание указателя на слайс (срез)
	slicePtr := new([]int)
	fmt.Println("Значение по умолчанию:", *slicePtr) // Выводит []

	// Создание указателя на структуру
	type Person struct {
		Name string
		Age  int
	}
	personPtr := new(Person)
	fmt.Println("Значение по умолчанию:", *personPtr) // Выводит {"" 0}
}
