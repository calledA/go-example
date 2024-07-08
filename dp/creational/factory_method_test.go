package creational

import "testing"

func TestFM(t *testing.T) {
	appleFac := new(AppleFactory)
	apple := appleFac.CreateFruit()
	apple.Show()

	bananaFac := new(BananaFactory)
	banana := bananaFac.CreateFruit()
	banana.Show()
}