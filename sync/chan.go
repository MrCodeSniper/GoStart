package main

import (
	"fmt"
)

func main() {
	demo1()
}


/**---------------创建一个结构包含数据载体和结果载体---------------------**/

type Worker struct {
	data chan int
	done chan bool
}


/**---------------核心工作---------------------**/
func doWorker(id int,data chan int,done chan bool){
	for n:=range data{
		fmt.Printf("worker id %d,received %c \n", id, n)
		done <- true
	}
}

/**---------------创建线程 返回通道---------------------**/
func createWorker(id int) Worker{ //返回只能入的通道
	c:=Worker{
		make(chan int),
		make(chan bool),
	}
	go doWorker(id,c.data,c.done)
	return c
}

func demo1(){
	var channels [10]Worker

	for i:=0;i<10;i++{
		channels[i]= createWorker(i)//创建对应的channel
	}

	for i,worker:= range channels{
		worker.data <- 'a'+i
	}

	/**---------------等待全部结束再继续执行---------------------**/
	for _,worker :=range channels{
		<-worker.done //收到通道发的值后 继续循环
	}

	for i, worker := range channels {
		worker.data <- 'A' + i
	}

	/**---------------等待全部结束---------------------**/
	for _,worker :=range channels{
		<-worker.done
	}

	/**---------------休眠一段时间保证其他线程执行完毕---------------------**/
	//time.Sleep(time.Second)








}