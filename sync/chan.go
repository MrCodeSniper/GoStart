package main

import (
	"GoStart/tree"
	"fmt"
	"sync"
)

func main() {
	demo3()
}


/**---------------创建一个结构包含数据载体和结果载体---------------------**/

type Worker struct {
	data chan int
	done chan bool
	//引用waitgroup帮助同步
	wg  *sync.WaitGroup
	//引用只执行函数进行提取
	fdone func()
}


/**---------------核心工作---------------------**/
func doWorker(id int,data chan int,done chan bool,fdone func()){
	for n:=range data{
		fmt.Printf("worker id %d,received %c \n", id, n)
		//done <- true
		//group.Done() //完成任务
		fdone()
	}
}

/**---------------创建线程 返回通道---------------------**/
func createWorker(id int,fdone func()) Worker{ //返回只能入的通道
	c:=Worker{
		make(chan int),
		make(chan bool),
		nil,
		fdone,
	}
	go doWorker(id,c.data,c.done,c.fdone)
	return c
}

func demo1(){
	var channels [10]Worker

	for i:=0;i<10;i++{
		channels[i]= createWorker(i,nil)//创建对应的channel
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
		<-worker.done //如果数据没有出来 就会堵塞 直到通道有数据
	}

	/**---------------休眠一段时间保证其他线程执行完毕---------------------**/
	//time.Sleep(time.Second)
}


/**---------------WaitGroup add添加任务 wait等待任务 done完成任务---------------------**/

func demo2(){


	var wg sync.WaitGroup //帮助同步

	var channels [10]Worker

	for i:=0;i<10;i++{
		channels[i]= createWorker(i,wg.Done)//创建对应的channel 传函数做参数
	}

	wg.Add(10)

	for i,worker:= range channels{
		worker.data <- 'a'+i
	}

	wg.Wait()

}

/**---------------用channel作为队列进行二叉树遍历---------------------**/

func demo3(){

	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	c:=root.TraverseByChannel()

	for node:= range c{
		fmt.Println(node.Value)
	}

}