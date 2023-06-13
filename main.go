package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/google/uuid"
)

var (
	DefaultPeople = People{
		Id:        uuid.New(),
		Name:      "Test",
		LastName:  "TestTest",
		Age:       18,
		CreatedAt: time.Now().Add(1 * time.Second),
	}
)

type People struct {
	Id        uuid.UUID
	Name      string
	LastName  string
	Age       int
	CreatedAt time.Time
}

func (p *People) copyDefaultValue() {
	pReflect := reflect.ValueOf(p).Elem()
	defaultReflect := reflect.ValueOf(DefaultPeople)

	for i := 0; i < pReflect.NumField(); i++ {
		if pReflect.Field(i).IsZero() {
			pReflect.Field(i).Set(defaultReflect.Field(i))
		}
	}
}

func main() {
	people := People{
		Name: "Lucas",
		Age:  30,
	}

	people.copyDefaultValue()

	fmt.Println(people)
}
