package structural

import "fmt"

type Goods struct {
	Kind string
	Fact bool
}

type Shopping interface {
	Buy(good *Goods)
}

type ChinaShopping struct{}

func (cs *ChinaShopping) Buy(good *Goods) {
	fmt.Println("在中国购物",good.Kind)
}

type Proxy struct {
	shopping Shopping
}

func (proxy *Proxy) Buy(good *Goods) {
	fmt.Println("proxy的代理方法")
	cs := new(ChinaShopping)
	cs.Buy(good)
}

func NewProxy(shopping Shopping) Shopping {
	return &Proxy{shopping}
}