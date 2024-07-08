package creational

import (
	"fmt"
	"sync"
	// "sync/atomic"
)

// 优点：
// (1) 单例模式提供了对唯一实例的受控访问。
// (2) 节约系统资源。由于在系统内存中只存在一个对象。    
// 缺点：
// (1) 扩展略难。单例模式中没有抽象层。
// (2) 单例类的职责过重。

type singleton struct{}

var instance *singleton = new(singleton)
var lock sync.Mutex
var initialized uint32
var once sync.Once

//饿汉式
// func GetInstance() *singleton {
// 	return instance
// }

//懒汉式,需要加锁，不然的话同时创造多个懒汉式实例会成为饿汉式
func GetInstance() *singleton {
	// if atomic.LoadUint32(&initialized) == 1 {
	// 	return instance
	// }

	// lock.Lock()
	// defer lock.Unlock()
	// if instance == nil {
	// 	instance = new(singleton)
	// 	atomic.StoreUint32(&initialized, 1)
	// }

	//sync.Once.Do()和保留内存的方式一样
	//func (o *Once) Do(f func()) {　　　//判断是否执行过该方法，如果执行过则不执行
	//   if atomic.LoadUint32(&o.done) == 1 {
	//       return
	//   }
	//   o.m.Lock()
	//   defer o.m.Unlock()　　
	//   if o.done == 0 {
	//       defer atomic.StoreUint32(&o.done, 1)
	//       f()
	//   }
	//}
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}

func (s *singleton) SomeThing() {
	fmt.Println("单例对象的方法")
}
