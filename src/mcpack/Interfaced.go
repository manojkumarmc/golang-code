package mcpack

import (
	"fmt"
)

type automobile interface {
	start()
	stop()
}

type car struct {
	make 	string
	model 	string
}

type bike struct {
	make 	string
	model	string
}

func (c car) start() {
	fmt.Println("Started the car " + c.make + " : " + c.model + ".")
}

func (c car) stop() {
	fmt.Println("Stopped the car " + c.make +  " : " + c.model + ".")
}

func carOperations(a automobile) {
	a.start()
	a.stop()
}

func Drive() {
	c := car{"ford", "figo"}
	carOperations(c)
}
