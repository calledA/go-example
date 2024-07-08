package structural

import "testing"

func TestDec(t *testing.T) {
	hw := new(HuaWei)
	hw.Show()

	moHW := NewMoDecorator(hw)
	moHW.Show()
}