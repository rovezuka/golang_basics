package main

import (
	"fmt"
	"math"
)

func main() {
	var todoList = [4]string{ // создание массива с количеством элементов и их типом
		"Купить хлеб",
		"Купить молоко",
		"Купить пиво",
	}
	fmt.Printf("Количество элементов в списке: %d\n", len(todoList)) // вывод длины массива

	todoList[3] = "Купить масло" // добавление элемента по индексу в массив

	for index, item := range todoList { // проход по массиву с индексем и значением
		fmt.Printf("%d. %s\n", index, item)
	}

	var infiniteList = [...]string{"Длина массива определяется количеством инициализаторов"}
	fmt.Printf("Количество элементов в списке: %d\n", len(infiniteList))

	for i := 0; i < 10; i++ { // обычная итерация цикла for
		fmt.Println(i)
		if i == 5 {
			break
		}
	}

	var arr [3]int // создание пустого массива
	fmt.Println(arr)
	arr = fillArray(arr)
	fmt.Println(arr)

	slices()

	var products = [3]string{"Хлеб", "Молоко", "Яблоко"}
	products = coup(products) // переворачиваем массив
	fmt.Println(products)

	throughNumber(10) // вывести квадраты 2 от единицы до десяти

	makeCopy() // функции make() и copy()
}

func fillArray(arr [3]int) [3]int { // функция заполнения массива значениями
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	return arr
}

func slices() { // срезы
	todoList := []string{ // создание среза
		"Купить хлеб",
		"Купить молоко",
		"Купить пиво",
	}
	fmt.Println("Длина среза:", len(todoList))
	fmt.Println("Емкость среза:", cap(todoList))

	for index, item := range todoList {
		fmt.Printf("%d. %s\n", index, item)
	}

	todoList = append(todoList, "Купить девушке цветы") // функция добавления элемента в срез
	fmt.Println("Длина среза:", len(todoList))
	fmt.Println("Емкость среза:", cap(todoList))

	var arr = [3]string{
		"Пройти курс по Go",
		"Сделать ДЗ по английскому",
		"Купить таблетки от головы",
	}

	// изменения в массиве отображаются на срезе, который на него ссылается
	var tasks = arr[:2]
	for i := range tasks {
		fmt.Println(tasks[i])
	}
	arr[0] = "Поспать дома и ничего не делать"
	for i := range tasks {
		fmt.Println(tasks[i])
	}
}

func coup(arr [3]string) [3]string { // функция для переворота массива
	var newArr [len(arr)]string
	x := 0
	for i := len(arr) - 1; i >= 0; i -= 1 {
		newArr[x] = arr[i]
		x++
	}
	return newArr
}

func throughNumber(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(math.Pow(2, float64(i)))
	}
}

func makeCopy() {
	// Создание слайса с длиной 5 и емкостью 10
	nums := make([]int, 5, 10)
	fmt.Println(nums) // Выводит [0 0 0 0 0]

	source := []int{1, 2, 3, 4, 5}
	destination := make([]int, len(source))

	numCopied := copy(destination, source)
	fmt.Println("Скопировано элементов:", numCopied)
	fmt.Println("Результат копирования:", destination)

	slice := []int{1, 2, 3, 4, 5}
	insert := []int{10, 20}

	// Вычисляем индекс вставки (2)
	index := 2

	// Расширяем слайс, чтобы вместить вставляемые элементы
	slice = append(slice, make([]int, len(insert))...)

	// Сдвигаем элементы вправо, начиная с индекса вставки
	copy(slice[index+len(insert):], slice[index:])

	// Вставляем элементы в слайс
	copy(slice[index:], insert)

	fmt.Println("Результат вставки:", slice)
}
