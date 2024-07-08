package structural

import "fmt"

type V5 interface {
	UserV5()
}

type AdapterPhone struct {
	v V5
}

func NewPhone(v V5) *AdapterPhone {
	return &AdapterPhone{v}
}

func (p *AdapterPhone) Charge() {
	fmt.Println("充电中")
	p.v.UserV5()
}

type V220 struct {}

func (v *V220) Use220V() {
	fmt.Println("使用220v")
}

type Adapter struct {
	v220 *V220
}

func (a *Adapter) UserV5() {
	fmt.Println("使用适配器")
	a.v220.Use220V()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220}
}