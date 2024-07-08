package behavioral

import "fmt"

type Beverage interface {
	BoilWater()
	Brew()
	PourInCup()
	AddThings()
	WantAddThings() bool
}

type template struct {
	b Beverage
}

func (t *template) MakeBerverage() {
	if t == nil {
		return
	}

	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()
	if t.b.WantAddThings() == true {
		t.b.AddThings()
	}
}

type MakeCoffee struct {
	template
}

func NewMakeCoffee() *MakeCoffee {
	makeCoffe := new(MakeCoffee)
	makeCoffe.b = makeCoffe
	return makeCoffe
}

func (mc *MakeCoffee) BoilWater() {
	fmt.Println("BoilWater")
}

func (mc *MakeCoffee) Brew() {
	fmt.Println("Brew")
}

func (mc *MakeCoffee) PourInCup() {
	fmt.Println("PourInCup")
}

func (mc *MakeCoffee) AddThings() {
	fmt.Println("AddThings")
}

func (mc *MakeCoffee) WantAddThings() bool {
	return true
}

