package creational

import "fmt"

type Fruit interface {
	Show()
}

type Apple struct {
	Fruit
}

func (apple *Apple) Show() {
	fmt.Println("苹果的show")
}

type Banana struct {
	Fruit
}

func (banana *Banana) Show() {
	fmt.Println("香蕉的show")
}

type Factory struct {}

func (factory *Factory) CreateFruit(kind string) Fruit {
	var fruit Fruit
	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "Banana" {
		fruit = new(Banana)
	}
	return fruit
}

