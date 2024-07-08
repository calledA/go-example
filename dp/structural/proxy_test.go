package structural

import "testing"

func TestProxy(t *testing.T) {
	g := Goods{
		Kind: "面膜",
		Fact: true,
	}
	var proxy,shopping Shopping
	proxy = NewProxy(shopping)
	proxy.Buy(&g)
}
