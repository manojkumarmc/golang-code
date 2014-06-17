package mcpack

import (
	"fmt"
)

type emp struct {
	id		int
	name	string
	age		int
}

func (e emp) getName() string {
	return e.name
}

func (e emp) setName(n string) {
	e.name = n
}

func StructProcesser() {
	e1 := emp{1, "Manoj Kumar MC", 36}
	fmt.Println(e1.getName())
}

