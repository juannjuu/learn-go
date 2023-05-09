package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required: "true" max: "10"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	sample := Sample{"Juan"}

	var sampleType reflect.Type = reflect.TypeOf(sample)
	fmt.Println(sampleType)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0))
	fmt.Println(sampleType.Field(0).Tag.Get("required")) //DOESNT WORK
	fmt.Println(sampleType.Field(0).Tag.Get("max"))      //DOESNT WORK

	/**
	Tag Required and Max is work but Properties Field Tag Get cannot read
	*/
	fmt.Println(IsValid(sample))

	sample.Name = ""
	fmt.Println(sample.Name)
	fmt.Println(IsValid(sample))

}
