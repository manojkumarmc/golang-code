package mcpack

import (
	"fmt"
)

type Person struct {
	Id		int
	Name	string
}

var people []Person

func Add(P Person) {
	people = append(people, P)
}

func List() {
	fmt.Println(people)
}

