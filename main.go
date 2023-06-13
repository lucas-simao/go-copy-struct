package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/google/uuid"
)

type People struct {
	Id        uuid.UUID
	Name      string
	LastName  string
	Age       int
	CreatedAt time.Time
}

func main() {
	id := uuid.New()

	people1 := People{
		Name: "Lucas",
		Age:  30,
	}

	defaultPeople := People{
		Id:        id,
		Name:      "Test",
		LastName:  "TestTest",
		Age:       18,
		CreatedAt: time.Now().Add(1 * time.Second),
	}

	// Initialize reflect.ValueOf with pointer to struct
	pReflect := reflect.ValueOf(&people1).Elem()
	defaultReflect := reflect.ValueOf(defaultPeople)

	for i := 0; i < pReflect.NumField(); i++ {
		// Verifica se o valor do campo é zero "sendo zero o valor inicial da variável"
		if pReflect.Field(i).IsZero() {
			// Caso seja zero, adiciona o valor default do outro field, dessa forma temos como criar um padrão para teste.
			pReflect.Field(i).Set(defaultReflect.Field(i))
		}
	}

	// Transforma o reflect.ValueOf em interface e faz assertion
	fmt.Printf("%+v\n", pReflect.Interface().(People))

	// Mas também podemos usar a struct usada para iniciar o reflectOf tendo em vista que pegamos a referencia da memoria na sua inicialização.
	fmt.Printf("%+v\n", people1)
}
