package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Текущее время:", currentTime)

	futureTime := currentTime.Add(time.Hour * 24)
	fmt.Println("Время через 24 часа:", futureTime)

	duration := futureTime.Sub(currentTime)
	fmt.Println("Прошло времени:", duration)

	formattedTime := currentTime.Format("02-Jan-2006 15:04:05")
	fmt.Println("Отформатированное время:", formattedTime)
}
