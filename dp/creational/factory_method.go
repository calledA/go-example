package creational

import "fmt"

type Fruit_1 interface {
	Show()
}

type AbstractMethodFactory interface {
	CreateFruit() Fruit_1
}

type Apple_1 struct {
	Fruit_1
}

func (apple *Apple_1) Show() {
	fmt.Println("苹果")
}

type Banana_1 struct {
	Fruit_1
}

func (banana *Banana_1) Show() {
	fmt.Println("香蕉")
}

type AppleFactory struct {
	AbstractMethodFactory
}

func (fac *AppleFactory) CreateFruit() Fruit_1 {
	fruit := new(Apple_1)
	return fruit
}

type BananaFactory struct {
	AbstractFactory
}

func (fac *BananaFactory) CreateFruit() Fruit_1 {
	fruit := new(Apple_1)
	return fruit
}
