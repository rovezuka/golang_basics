package main

import "fmt"

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

}
