package main

import "fmt"

type employee struct { // создание структуры
	name   string
	sex    string
	age    int
	salary int
}

func newEmployee(name, sex string, age, salary int) employee { // функция создания экземпляра
	return employee{
		name:   name,
		sex:    sex,
		age:    age,
		salary: salary,
	}
}

func (em employee) getInfo() string { // получить информацию об экземпляре
	return fmt.Sprintf("Сотрудник: %s\nВозраст: %d\nЗарплата: %d\n", em.name, em.age, em.salary)
}

// без звездочки - копия объекта, со звездочкой сам объект(ресивер)
func (em *employee) setName(name string) { // изменить имя (метод)
	em.name = name
}

func main() {
	employee1 := newEmployee("Ростислав", "М", 19, 10000)
	fmt.Printf("%s\n", employee1.getInfo())
	employee1.setName("Ярослав")
	fmt.Printf("%s\n", employee1.getInfo())
}
