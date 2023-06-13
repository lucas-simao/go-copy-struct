package main

import (
	"reflect"
	"testing"
	"time"
)

func TestPeople(t *testing.T) {
	createdAt := time.Now()

	tests := map[string]struct {
		expect People
		people People
	}{
		"With name and createdAt": {
			people: People{
				Name:      "Test",
				CreatedAt: createdAt,
			},
			expect: func() People {
				p := DefaultPeople
				p.Name = "Test"
				p.CreatedAt = createdAt
				return p
			}(),
		},
		"All fields empty": {
			people: People{},
			expect: DefaultPeople,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.people.copyDefaultValue()

			p := reflect.ValueOf(test.people)
			pExpected := reflect.ValueOf(test.expect)

			for i := 0; i < p.NumField(); i++ {
				if !p.Field(i).Equal(pExpected.Field(i)) {
					t.Errorf("Field: %s expect: %v got: %v", p.Type().Field(i).Name,
						pExpected.Field(i), p.Field(i),
					)
				}
			}
		})
	}
}
