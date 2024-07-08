package behavioral

import "testing"

func TestXxx(t *testing.T) {
	hero := Hero{}

	hero.SetWeaponStrategy(new(AK))
	hero.Fight()
}
