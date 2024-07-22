package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	Name string
}

var instance *Singleton
var once sync.Once

func GetInstance(name string) *Singleton {
	once.Do(func() {
		instance = &Singleton{
			Name: name,
		}
	})
	return instance
}

func main() {
	instance = GetInstance("golang")
	instance = GetInstance("java")
	fmt.Println(instance.Name == "golang")
	fmt.Println(instance.Name != "java")
}
