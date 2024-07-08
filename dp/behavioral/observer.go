package behavioral

import "fmt"

type Listener interface {
	OnTeacherComming()
}

type Notifier interface {
	AddListener(listener Listener)
	RemoveListener(listener Listener)
	Notify()
}

type StuOne struct {
	Badthing string
}

func (s *StuOne) OnTeacherComming() {
	fmt.Println("StuOne stop",s.Badthing)
}

func (s *StuOne) DoBadthing() {
	fmt.Println("StuOne do",s.Badthing)
}

type StuTwo struct {
	Badthing string
}

func (s *StuTwo) OnTeacherComming() {
	fmt.Println("StuTwo stop",s.Badthing)
}

func (s *StuTwo) DoBadthing() {
	fmt.Println("StuTwo do",s.Badthing)
}

type Monitor struct {
	ListenerList []Listener
}

func (monitor *Monitor) AddListener(listener Listener) {
	monitor.ListenerList = append(monitor.ListenerList, listener)
}

func (monitor *Monitor) RemoveListener(listener Listener) {
	for index, l := range monitor.ListenerList {
		//找到要删除的元素位置
		if listener == l {
			//将删除的点前后的元素链接起来
			monitor.ListenerList = append(monitor.ListenerList[:index], monitor.ListenerList[index+1:]...)
			break
		}
	}
}

func (monitor *Monitor) Notify() {
	for _, listener := range monitor.ListenerList {
		listener.OnTeacherComming()
	}
}