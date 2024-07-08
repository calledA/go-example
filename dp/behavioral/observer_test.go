package behavioral

import "testing"

func TestObserver(t *testing.T) {
	s1 := &StuOne{
		Badthing:"抄作业",
	}

	s2 := &StuTwo{
		Badthing:"玩游戏",
	}

	monitor := new(Monitor)

	s1.DoBadthing()
	s2.DoBadthing()

	monitor.AddListener(s1)
	monitor.AddListener(s2)

	monitor.Notify()
}