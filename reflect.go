package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	// tag required dan max
	Name string `required:"true" max:"10"`
}

type Person struct {
	// tag required dan max
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		// ambil nama field
		fmt.Println(structField.Name, "with type", structField.Type)
		// ambil tag
		fmt.Println(structField.Tag.Get("required"))
		// ambil tag max
		fmt.Println(structField.Tag.Get("max"))
	}
}

func IsValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			// jika result false maka return result
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	// read field
	readField(Sample{"Aidil"})
	// read field person
	readField(Person{"Aidil", "", ""})
	// validasi data
	person := Person{
		Name:    "ada",
		Address: "ada",
		Email:   "ada",
	}
	fmt.Println(IsValid(person))
}
