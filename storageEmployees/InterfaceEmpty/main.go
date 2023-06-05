package main

import "fmt"

func main() {
	printType(3)
	printType("string")
	printType([]string{"слайс"})
}

func printType(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("тип аргумента int")
	case string:
		fmt.Println("тип аргумента string")
	default:
		fmt.Println("тип аргумента не int и не string")
	}
}
