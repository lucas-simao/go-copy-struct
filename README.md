# go-copy-struct
Using go to validate if the structs have a empty value and than copy a default value.

## This is a example how to copy only when empty value from one default struct to another struct.

```go
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

	// Initialize reflect.ValueOf with pointer to struct, this is used to manipulate values
	pReflect := reflect.ValueOf(&people1).Elem()
	defaultReflect := reflect.ValueOf(defaultPeople)

	for i := 0; i < pReflect.NumField(); i++ {
		// Check if field is empty value
		if pReflect.Field(i).IsZero() {
			// If is empty, copy value from default struct 
			pReflect.Field(i).Set(defaultReflect.Field(i))
		}
	}

	// To transform reflect.ValueOf to interface and make assertion
	fmt.Printf("%+v\n", pReflect.Interface().(People))

	// You can access the values from the struct used with a pointer too.
	fmt.Printf("%+v\n", people1)
}

```
