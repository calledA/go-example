package behavioral

import "testing"

func TestTemplate(t *testing.T) {
	makeCoffee := NewMakeCoffee()
	makeCoffee.MakeBerverage()
}