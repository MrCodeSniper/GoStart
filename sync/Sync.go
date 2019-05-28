package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
 操作系统本身可以进行线程和进程的调度，具备并发处理能力
 GO在用户层上再构建了一级调度 将并发的粒度降低 更大限度提升程序运行效率

 通过go 函数 启动一个goroutine




 */





//单独创建goroutine
func main() {

	go func() {
		sum:=0
		for i:=0;i<100000;i++{
			sum+=i
		}
		println(sum)
		time.Sleep(10*time.Second)
	}()


   fmt.Println(runtime.NumGoroutine())

   time.Sleep(5*time.Second)

}
