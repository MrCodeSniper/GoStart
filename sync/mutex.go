package main

import (
	"fmt"
	"sync"
	"time"
)

//go run mutex.go
//go run -race mutex.go 查看冲突
func main() {
	interrupt()
}

type AtomincInt struct {
	i int
	lock sync.Mutex
}

func (i *AtomincInt) increment(){
	func() { //限定同步在目前代码块
		i.lock.Lock()
		defer i.lock.Unlock()
		(*i).i++
	}()
}

func (i *AtomincInt) get() int{
	i.lock.Lock()
	defer  i.lock.Unlock()
	return int((*i).i)
}


func interrupt(){ //加了锁以后没有datarace
	var v AtomincInt

	v.increment()

	go func() {
		v.increment()
	}()

	/**---------------多线程写的情况会出现同时修改数据 导致冲突---------------------**/

	time.Sleep(time.Second)

	fmt.Println(v.get())
}
