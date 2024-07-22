package main

import "fmt"

func main() {
	dogFactory := &DogFactory{}
	catFactory := &CatFactory{}

	dog := dogFactory.New()
	cat := catFactory.New()

	dog.Says()
	cat.Says()

	fmt.Println("Dog likes water? ", dog.LikesWater())
	fmt.Println("Cat likes water? ", cat.LikesWater())
}
