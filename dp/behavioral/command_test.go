package behavioral

import "testing"

func TestCmd(t *testing.T) {
	doctor := new(Doctor)
	cmdEye := CommandEye{doctor }
	cmdEye.Treat()

	CmdNose := CommandNose{doctor }
	CmdNose.Treat()
}