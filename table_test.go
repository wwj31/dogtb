package dogtb

import (
	"fmt"
	"testing"
)

func TestCreateArrayStruct(t *testing.T) {
	type St struct {
		Name  string
		Age   int `dogtg:"user's Age'"`
		Class int
	}

	st1 := &St{
		Name:  "530µs",
		Age:   15,
		Class: 1,
	}
	st2 := &St{
		Name:  "nibgy",
		Age:   100,
		Class: 2,
	}
	st3 := &St{
		Name:  "zhgyru",
		Age:   31,
		Class: 10000000000,
	}

	array := append([]*St{}, st1, st2, st3)

	tab, _ := Create(array)
	fmt.Println(tab.String())
	fmt.Println(tab.String(1))
	fmt.Println(tab.String(2))
	fmt.Println(tab.String(3))
}
