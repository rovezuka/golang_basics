package main // Объявление название пакета

import "fmt" //  Импорт пакета fmt из стандартной библиотеки

const pi = 3.145 // глобальная + константа. Неизменяемая переменная

func main() { // Объявление функции main() которая является точкой входа для любой программы
	fmt.Println("Hello, world!")
	fmt.Println(pi)
	var phrase = "Rostislav, hello!"
	fmt.Println(phrase)
	var randomNumber int = 11
	fmt.Println("My number in VityazGTU is", randomNumber)
	localVar := "bibabobaflex" // локальная переменная
	fmt.Println(localVar)
	var a, b, c int
	a, b, c = 1, 2, 3 // множественное присваивание
	fmt.Println(a, b, c)

	var golang = "Go рулит!" // на этапе компиляции, тип переменной определяется как string автоматически
	fmt.Println(golang)
	var h, f, s = true, 2.3, "four" // bool, float64, string, компилятор автоматически определяет тип
	fmt.Println(h, f, s)

	circleRadius := 2
	fmt.Printf("Радиус круга: %d см.\n", circleRadius)
	// для мат. операций необходимо привести переменные к одному типу
	circleArea := float32(circleRadius) * float32(circleRadius) * pi
	fmt.Printf("Площадь круга: %f см. кв.", circleArea)
}
