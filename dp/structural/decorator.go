package structural

import "fmt"

type Phone interface {
	Show()
}

type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {

}

type HuaWei struct{}

func (hw *HuaWei) Show() {
	fmt.Println("show huawei")
}

type MoDecorator struct {
	Decorator
}

func (md *MoDecorator) Show() {
	md.phone.Show()
	fmt.Println("贴膜的手机")
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone}}
}

