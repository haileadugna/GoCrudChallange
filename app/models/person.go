package models

import (
	"sync"
	"github.com/google/uuid"
)

type Person struct {
	ID      string   `json:"id"`
	Name    string   `json:"name" binding:"required"`
	Age     int      `json:"age" binding:"required"`
	Hobbies []string `json:"hobbies"`
}

var (
	persons = make(map[string]Person)
	mutex   sync.RWMutex
)

func GetAllPersons() []Person {
	mutex.RLock()
	defer mutex.RUnlock()

	var result []Person
	for _, person := range persons {
		result = append(result, person)
	}
	return result
}

func GetPerson(id string) (Person, bool) {
	mutex.RLock()
	defer mutex.RUnlock()

	person, exists := persons[id]
	return person, exists
}

func CreatePerson(person Person) (Person, bool) {
	mutex.Lock()
	defer mutex.Unlock()

	person.ID = uuid.New().String()
	persons[person.ID] = person
	return person, true
}

func UpdatePerson(id string, person Person) (Person, bool) {
	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := persons[id]; !exists {
		return Person{}, false
	}

	person.ID = id
	persons[id] = person
	return person, true
}

func DeletePerson(id string) (Person, bool) {
	mutex.Lock()
	defer mutex.Unlock()

	person, exists := persons[id]
	if !exists {
		return Person{}, false
	}

	delete(persons, id)
	return person, true
}
