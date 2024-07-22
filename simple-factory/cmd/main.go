package main

import (
	"fmt"
	"github.com/iarsham/design-patterns/simple-factory/models"
	"github.com/iarsham/design-patterns/simple-factory/service"
)

func main() {
	factory := service.New("sms", "+1 123 456 7890")
	factory.Send("Hello World!")

	factory = service.New("email", "RqXp9@example.com")
	factory.Send("Hello World!")

	product := models.Product{}
	productFactory := product.New("Product 1")
	fmt.Println("My product was created at: ", productFactory.CreatedAt.UTC())
}
