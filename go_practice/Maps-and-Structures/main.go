package main

import (
	"fmt"
	"math"
)

type age int // кастомный тип

type Person struct { // создание структуры
	name string
	age  int
}

type Employee struct {
	Person
	salary float64
}

type Address struct {
	Street     string
	City       string
	PostalCode string
}

type Man struct { // вложенные структуры
	Name    string
	Age     int
	Address Address
}

type Circle struct { // круг
	radius float64
}

func main() {
	var mapsOne map[string]int // создание мапы
	fmt.Println(mapsOne)

	ages := make(map[string]int) // создание мапы через make
	fmt.Println(ages)

	ages["Максим"] = 22 // присваиваем значения
	ages["Александр"] = 25
	ages["Ярослав"] = 26

	for key, value := range ages {
		fmt.Printf("Возраст %sа -> %d\n", key, value)
	}

	ages["Антон"] = 25

	age, exists := ages["Антон"] // 25, true
	if exists {
		fmt.Printf("Возраст Антона: %d\n", age)
	} else {
		fmt.Println("Антона нет в списке")
	}

	delete(ages, "Антон") // удаление значения
	fmt.Println(ages)

	mapsTwo := map[string]int{ // задать значения мапы первоначально
		"Один": 1,
		"Два":  2,
		"Три":  3,
	}
	fmt.Println(mapsTwo)

	customTypes()

	man := Person{ // создание экземпляра
		name: "Rostislav",
		age:  19,
	}
	man.sayHello() // вызов метода структуры

	emp := Employee{
		Person: Person{name: "John"}, // встраивание
		salary: 5000,
	}

	emp.sayHello() // Метод sayHello() доступен через встраивание

	// реализация вложенности структур
	person := Man{
		Name: "John Doe",
		Age:  30,
		Address: Address{
			Street:     "123 Main St",
			City:       "New York",
			PostalCode: "10001",
		},
	}

	fmt.Println(person.Name)               // Выводит: John Doe
	fmt.Println(person.Address.Street)     // Выводит: 123 Main St
	fmt.Println(person.Address.City)       // Выводит: New York
	fmt.Println(person.Address.PostalCode) // Выводит: 10001

	circleArea := Circle{
		radius: 5,
	}
	fmt.Println(circleArea.calculateArea())

}

func (a age) isAdult() bool { // метод кастомного типа
	return a >= 18
}

func customTypes() {
	myAge := age(22)

	fmt.Println("Я совершеннолетний?:", myAge.isAdult())
}

func (p Person) sayHello() {
	fmt.Println("Hello, my name is", p.name)
}

func (c Circle) calculateArea() float64 { // метод вычисления площади круга
	return math.Pi * math.Pow(c.radius, 2)
}
