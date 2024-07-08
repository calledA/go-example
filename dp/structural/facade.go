package structural

import "fmt"

type SubSystemA struct{}

func (sa *SubSystemA) MethodA() {
	fmt.Println("子系统方法A")
}

type SubSystemB struct {}

func (sb *SubSystemB) MethodB() {
	fmt.Println("子系统方法B")
}


type SubSystemC struct {}

func (sb *SubSystemC) MethodC() {
	fmt.Println("子系统方法C")
}

type Facade struct {
	a *SubSystemA
	b *SubSystemB
	c *SubSystemC
}

func (f *Facade) MethodOne() {
	f.a.MethodA()
	f.b.MethodB()
}

func (f *Facade) MethodTwo() {
	f.c.MethodC()
}