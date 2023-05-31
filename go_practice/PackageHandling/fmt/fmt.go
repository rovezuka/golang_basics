package main

import (
	"fmt"
	"os"
)

func main() {
	number := 3.14159
	fmt.Printf("Число: %.5f\n", number) // количество знаков после запятой, %.5f - 5 цифр

	/*Пакет fmt позволяет выводить данные не только на стандартный вывод, но и на произвольные потоки,
	такие как файлы или сетевые соединения.*/
	file, err := os.Create("output.txt") // создание файла
	if err != nil {                      // проверка на создание файлы
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer file.Close() //

	fmt.Fprintln(file, "Это текст, записанный в файл")

	/* При использовании функций Scan и Scanln для чтения данных из стандартного ввода,
	необходимо проверять ошибки.*/
	var age int

	fmt.Print("Введите ваш возраст: ")
	_, err2 := fmt.Scanln(&age)
	if err2 != nil { // проверка, что введен integer
		fmt.Println("Ошибка чтения ввода:", err2)
		return // завершение работы функции main
	}

	fmt.Println("Ваш возраст:", age)
}
