package creational

import "fmt"

type Programmer interface {
	Work()
}

type Architest interface {
	Design()
}

type FrontArchitest struct{}

func (front *FrontArchitest) Design() {
	fmt.Println("前端架构师的设计")
}

type FrontProgrammer struct {}

func (front *FrontProgrammer) Work() {
	fmt.Println("前端程序员的工作")
}

type AbstractFactory interface {
	CreateArchitest() Architest
	CreateProgrammer() Programmer
}

type FrontFactory struct {}

func (front *FrontFactory) CreateProgrammer() Programmer {
	return &FrontProgrammer{}
}

func (front *FrontFactory) CreateArchitest() Architest {
	return &FrontArchitest{}
}


