package creational

import "testing"

func TestSingle(t *testing.T) {
	s := GetInstance()
	s.SomeThing()
}