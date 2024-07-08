package behavioral

import "fmt"

type Doctor struct{}

func (d *Doctor) treatEye() {
	fmt.Println("治疗眼睛")
}

func (d *Doctor) treatNose() {
	fmt.Println("治疗鼻子")
}

type Command interface {
	Treat()
}

type CommandEye struct {
	doctor *Doctor
}

func (cmd *CommandEye) Treat() {
	cmd.doctor.treatEye()
}

type CommandNose struct {
	doctor *Doctor
}

func (cmd *CommandNose) Treat() {
	cmd.doctor.treatNose()
}