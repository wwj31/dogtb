package dogtb

import (
	"fmt"
	"testing"
)

func TestCreateArrayStruct(t *testing.T) {
	type St struct {
		Name  string
		Id    string
		Age   int `tb:"user's Age"`
		Class int
	}

	array := []*St{
		{Name: "wwj", Id: "502948676638566431", Age: 15, Class: 1},
		{Name: "bartholomew", Id: "501925768954674307", Age: 100, Class: 2},
		{Name: "bobo", Id: "52258675940625906", Age: 31, Class: 3},
	}
	tab, _ := Create(array)

	fmt.Println(tab.String())  // style 0
	fmt.Println(tab.String(1)) // style 1
	fmt.Println(tab.String(2)) // style 2
	fmt.Println(tab.String(3)) // style 3
}
