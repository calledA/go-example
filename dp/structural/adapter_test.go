package structural

import "testing"

func TestAdapter(t *testing.T) {
 iphone := NewPhone(NewAdapter(new(V220)))
 iphone.Charge()
}