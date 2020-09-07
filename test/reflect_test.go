package test

import (
	"fmt"
	"reflect"
	"testing"
)

func GetTag(i interface{}) int {
	s := reflect.TypeOf(i).Elem()

	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		fmt.Println(field.Name)
		fmt.Println(field.Tag.Get("protobuf")) //将tag输出出来
	}

	return 0
}

type data struct {
	Name string
	Age  int
}

func TestGetTage(t *testing.T) {
	GetTag(&data{})
}
