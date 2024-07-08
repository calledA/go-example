package creational

import "testing"

func TestSF(t *testing.T) {
	fac := new(Factory)
	apple := fac.CreateFruit("apple")
	apple.Show()
}