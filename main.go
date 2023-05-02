package main

import "fmt"

func main() {
	// вывод текста через функцию Println из пакета "fmt"
	/* Просто
	многострочный
	комментарий
	*/
	const pi = 3.14
	fmt.Println("Hello, world!")
	x, y := 18, 19
	fmt.Println(x, y)
	fmt.Println(pi)
	var a int = 31
	var b int = 14
	var result int

	// сложение
	result = a + b
	fmt.Println(result) // выведет 45

	// вычитание
	result = a - b
	fmt.Println(result) // выведет 17

	// умножение
	result = a * b
	fmt.Println(result) // выведет 434

	// деление
	result = a / b
	fmt.Println(result) // выведет 2

	// остаток от деления
	result = a % b
	fmt.Println(result) // выведет 3
}
