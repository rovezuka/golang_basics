package main

import "fmt"

func main() {
	/* Просто
	многострочный
	комментарий
	*/

	/*
		// контстанта
		const pi = 3.14
		fmt.Println(pi)
		// вывод текста через функцию Println из пакета "fmt"
		fmt.Println("Hello, world!")

		// объявление переменных
		x, y := 18, 19
		fmt.Println(x, y)

		// арифмитические операции
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
		c := "Привет, "
		d := "Мир!"
		fmt.Println(c + d) // выведет: Привет, Мир!

		// операторы присваивания
		a += b
		fmt.Println(a) // выведет 45

		// инкремент и декремент
		a++ // аналог выражения a = a + 1
		fmt.Println(a) // выведет 3
		a-- // аналог выражения a = a - 1
		fmt.Println(a) // выведет 1

		// операторы сравнения и логические операторы
		// равно
		fmt.Println(a == b) // выведет false

		// не равно
		fmt.Println(a != b) // выведет true

		// больше чем
		fmt.Println(a > b) // выведет true

		// логическое И
		fmt.Println(a != b && a >= b) // выведет true

		// логическое ИЛИ
		fmt.Println(a == b || a > b) // выведет true

		// логическое НЕ
		fmt.Println(!(a < b)) // выведет true

		// ввод
		// вывод суммы чисел
		var j int
		var k int
		fmt.Scan(&j)
		fmt.Scan(&k)
		fmt.Println(j + k)

		// вывод через Scanln

		var j, k int
		fmt.Scanln(&j, &k)
		fmt.Println(j + k)


		// условный оператор if/else
		var pass int
		pass = 7055
		if pass == 7055 {
			fmt.Println("Пароль верный")
		}
		// пароль верный, условие равно true
		// код будет выволнен и выведет: Пароль верный

		pass = 7055
		if pass == 3142 {
			fmt.Println("Пароль неверный")
		}
		// пароль неверный, условие равно false
		// код не будет выполнен и ничего не будет выведено

		// else
		iq := 70

		if iq > 120 {
			fmt.Println("Вы гений!")
		} else {
			fmt.Println("Увы, но вы не гений")
		}
	*/

	// else if
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	if a > b {
		fmt.Println("a больше чем b")
	} else if a == b {
		fmt.Println("a равно b")
	} else if c := 100 - 1; c == 100 {
		fmt.Println("проверка работы условий")
	} else {
		fmt.Println("a меньше чем b")
	}
}
