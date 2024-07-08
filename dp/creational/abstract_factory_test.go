package creational

import "testing"

func TestAF(t *testing.T) {
	ff := new(FrontFactory)
	fa := ff.CreateArchitest()
	fp := ff.CreateProgrammer()
	fa.Design()
	fp.Work()
}