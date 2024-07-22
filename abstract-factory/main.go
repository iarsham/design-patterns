package main

func main() {
	woodenDoorFactory := &WoodenDoorFactory{}
	woodenDoor := woodenDoorFactory.MakeDoor()
	woodenDoorExpert := woodenDoorFactory.MakeExpert()
	woodenDoor.Description()
	woodenDoorExpert.Description()

	ironDoorFactory := &IronDoorFactory{}
	ironDoor := ironDoorFactory.MakeDoor()
	ironDoorExpert := ironDoorFactory.MakeExpert()
	ironDoor.Description()
	ironDoorExpert.Description()
}
