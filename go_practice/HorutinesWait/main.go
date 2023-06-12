package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go printMessage("Hello", &wg)
	go printMessage("Go", &wg)

	wg.Wait()

	fmt.Println("All messages printed")
}

func printMessage(message string, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Second)
	fmt.Println(message)
}
