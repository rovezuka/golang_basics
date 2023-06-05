package storage

import (
	"errors"
)

type Employee struct {
	id     int
	name   string
	age    string
	salary int
}

type Storage interface { // создание интерфейса и задание его методов
	insert(e Employee) error
	get(id int) (Employee, error)
	delete(id int) error
}

type MemoryStorage struct { // создание первой структуры
	data map[int]Employee
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[int]Employee),
	}
}

func (s *MemoryStorage) insert(e Employee) error {
	s.data[e.id] = e

	return nil
}

func (s *MemoryStorage) get(id int) (Employee, error) {
	e, exists := s.data[id]
	if !exists {
		return Employee{}, errors.New("Employee with such id doesn't exists")
	}

	return e, nil
}

func (s *MemoryStorage) delete(id int) error {
	delete(s.data, id)
	return nil
}
