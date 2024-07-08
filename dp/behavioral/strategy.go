package behavioral

import "fmt"

type WeaponStrategy interface {
	UseWeapon()
}

type AK struct{}

func (ak *AK) UseWeapon() {
	fmt.Println("AK的UseWeapon()")
}

type Kinfe struct {}

func (kinfe *Kinfe) UseWeapon() {
	fmt.Println("Kinfe的UseWeapon()")
}

type Hero struct {
	strategy WeaponStrategy
}

func (h *Hero) SetWeaponStrategy(s WeaponStrategy) {
	h.strategy = s
}

func (h *Hero) Fight() {
	h.strategy.UseWeapon()
}