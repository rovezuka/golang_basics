package main

import (
	"fmt"
	"time"
)

/* Переключением между горутинами занимается планировщик Go. Переключение
происходит при блокировке выполнения текущей горутины */

func main() { // главная функция также горутина, поэтому в конце нужен time.Sleep
	t := time.Now() // время старта

	fmt.Printf("Старт: %s\n", t.Format(time.RFC3339))

	go func() { // спиннер ожидания
		for {
			for _, r := range `\|/` {
				fmt.Printf("\r%c ", r)
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()

	channel1 := make(chan int)
	channel2 := make(chan int)

	// вычисления
	go calculateSomething(1000, channel1)

	// снова вычисления
	go calculateSomething(2000, channel2)

	// блокирует выполнение главной горутины, планировщик передает конекст на другие незаблокированные горутины
	// time.Sleep(7 * time.Second)

	// блокировка выполнения главной горутины не требуется, т. к. ее блокирует чтение из канала

	fmt.Println(<-channel1) // чтение из каналов
	fmt.Println(<-channel2)

	fmt.Printf("Время выполнения программы %s\n", time.Since(t))

	numbers := make(chan int)

	go generationNumbers(1000, numbers)

	for n := range numbers {
		fmt.Println(n)
	}

	number := make(chan int)

	select { // позволяет избежать deadlock, если канал пуст, не блокируя функцию main
	case n := <-number:
		fmt.Println(n)
	default:
		fmt.Println("В канале ничего нет.")
	}

	factN := factorial(45) // рекурсивный факториал 45
	fmt.Println(factN)

}

func calculateSomething(n int, res chan int) {
	t := time.Now()

	result := 0
	for i := 0; i <= n; i++ {
		result += i * 2
		time.Sleep(time.Millisecond * 3)
	}

	fmt.Printf("Время выполнения расчетов: %s\n", time.Since(t))
	res <- result // запись результата в канал
}

func generationNumbers(n int, res chan int) {
	for i := 0; i <= n; i++ {
		res <- i * 2
	}

	close(res) // закрытия канала, чтобы не произошел deadlock
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
