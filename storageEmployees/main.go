package main

import (
	"errors"
	"fmt"
)

type dumbStorage struct{} // создание второй структуры

type employee struct {
	id     int
	name   string
	age    string
	salary int
}

type storage interface { // создание интерфейса и задание его методов
	insert(e employee) error
	get(id int) (employee, error)
	delete(id int) error
}

type memoryStorage struct { // создание первой структуры
	data map[int]employee
}

func newDumbStorage() *dumbStorage {
	return &dumbStorage{}
}

func (s *dumbStorage) insert(e employee) error {
	fmt.Printf("вставка пользоватпеля с id: %d прошла успешно\n", e.id)
	return nil
}

func (s *dumbStorage) get(id int) (employee, error) {
	e := employee{
		id: id,
	}

	return e, nil
}

func (s *dumbStorage) delete(id int) error {
	fmt.Printf("удаления пользователя с id: %d прошла успешно\n", id)
	return nil
}

func newMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[int]employee),
	}
}

func (s *memoryStorage) insert(e employee) error {
	s.data[e.id] = e

	return nil
}

func (s *memoryStorage) get(id int) (employee, error) {
	e, exists := s.data[id]
	if !exists {
		return employee{}, errors.New("employee with such id doesn't exists")
	}

	return e, nil
}

func (s *memoryStorage) delete(id int) error {
	delete(s.data, id)
	return nil
}

func main() {
	var s storage

	fmt.Println("s == nil", s == nil)
	fmt.Printf("type of s: %T\n", s)

	s = newMemoryStorage()

	fmt.Println("s == nil", s == nil)
	fmt.Printf("type of s: %T\n", s)

	s = newDumbStorage()
	fmt.Println("s == nil", s == nil)
	fmt.Printf("type of s: %T\n", s)

	s = nil
	fmt.Println("s:", s)
	fmt.Printf("type of s: %T\n", s)

	// полиморфизм
	//создание двух структур, которые реализовывают один и тот же интерфейс
	ds := newDumbStorage()

	spawnEmployees(ds)

	ms := newMemoryStorage()

	spawnEmployees(ms)

	fmt.Println(ms.data)
}

func spawnEmployees(s storage) {
	for i := 1; i <= 10; i++ {
		s.insert(employee{id: i})
	}
}
