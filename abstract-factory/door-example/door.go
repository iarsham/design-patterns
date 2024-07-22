package main

import "fmt"

type Door interface {
	Description()
}

// WoodenDoor is a door made of wood
type WoodenDoor struct{}

func (w *WoodenDoor) Description() {
	fmt.Println("iam a wooden door")
}

// IronDoor is a door made of iron
type IronDoor struct{}

func (i *IronDoor) Description() {
	fmt.Println("iam an iron door")
}

type Expert interface {
	Description()
}

// Carpenter is an expert for build wooden doors
type Carpenter struct{}

func (c *Carpenter) Description() {
	fmt.Println("I can only fit wooden doors")
}

// Welder is an expert for build iron doors
type Welder struct{}

func (w *Welder) Description() {
	fmt.Println("I can only fit iron doors")
}

type DoorFactory interface {
	MakeDoor() Door
	MakeExpert() Expert
}

type WoodenDoorFactory struct{}

func (w *WoodenDoorFactory) MakeDoor() Door {
	return &WoodenDoor{}
}

func (w *WoodenDoorFactory) MakeExpert() Expert {
	return &Carpenter{}
}

type IronDoorFactory struct{}

func (i *IronDoorFactory) MakeDoor() Door {
	return &IronDoor{}
}

func (i *IronDoorFactory) MakeExpert() Expert {
	return &Welder{}
}

//
