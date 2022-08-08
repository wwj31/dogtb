package main

import (
	"fmt"
	"github.com/wwj31/dogtb/table"
)

type St struct {
	Name string
	Age  int `dogtg:"user's Age'"`
}

func main() {
	t, _ := table.Create(&St{
		Name: "wwj",
		Age:  99,
	})
	fmt.Println(t.String())
}
