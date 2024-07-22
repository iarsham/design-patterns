package main

import "fmt"

type Animal interface {
	Says()
	LikesWater() bool
}

type Dog struct{}

func (d *Dog) Says() {
	fmt.Println("woof")
}

func (d *Dog) LikesWater() bool {
	return true
}

type Cat struct{}

func (c *Cat) Says() {
	fmt.Println("meow")
}

func (c *Cat) LikesWater() bool {
	return false
}

type AnimalFactory interface {
	New() Animal
}

type DogFactory struct{}

func (d *DogFactory) New() Animal {
	return &Dog{}
}

type CatFactory struct{}

func (c *CatFactory) New() Animal {
	return &Cat{}
}
