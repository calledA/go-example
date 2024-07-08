package structural

import "testing"

func TestFacade(t *testing.T) {
	sa := new(SubSystemA)
	sa.MethodA()

	sb := new(SubSystemB)
	sb.MethodB()

	//外观模式
	f := Facade{
		a:new(SubSystemA),
		b:new(SubSystemB),
		c:new(SubSystemC),
	}
	f.MethodOne()
}