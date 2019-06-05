package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ( //包下变量
	variable1 int
	variable2 string
)

func main() {

	var a, b, c, d, e, f, g int //多值 定义赋值

	var i, j, k = 0, true, "xxx" //自动类型推断

	h := 0 //短变量声明 和赋值 并自动推断类型

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k)

	setValue()

	fmt.Println("Ranges")

	ranges()

	selectInChannel()

}

/**

 对左侧操作数中的表达式索引值(数组)进行计算和确定 然后赋值
如果右侧表达式计算引用左侧变量 则创建临时变量拷贝并完成

*/
func setValue() {
	x := []int{1, 2, 3}
	i := 0
	//先确定数组X[0]=2
	i, x[i] = 1, 2
	fmt.Println(i, x) //1,123 错误 TODO 1,223

	x = []int{1, 2, 3}
	i = 0
	x[i], i = 2, 1
	fmt.Println(i, x) //1,223

	x = []int{1, 2, 3}
	i = 0
	x[i], i = 2, x[i] //右边的x[i]拷贝原来的1,2,3
	fmt.Println(i, x) // 1,223

	x[0], x[0] = 1, 2
	fmt.Println(x[0]) //2

}

func ranges() {

	//sync.WaitGroup内部有一个计数器 有三个方法：Add(), Done(), Wait() 用来控制计数器的数量
	//Add(n) 把计数器设置为n ，Done() 每次把计数器-1 ，wait() 会阻塞代码的运行，直到计数器地值减为0。
	wg := sync.WaitGroup{}

	si := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range si { //i为数组下标
		wg.Add(1)
		//以并发的方式调用匿名函数func
		go func() {
			println(i)
			wg.Done()
		}()
	}

	wg.Wait()

}

/**---------------Select语句 任务分发 调度任务---------------------**/

func selectInChannel() {

	var c1, c2 chan int

	c1, c2 = generator(), generator()

	// 10s计时器，如果写在循环中会无限重置
	tm := time.After(10 * time.Second) //返回通道 到达10S发一条消息

	// 1s计时器
	tick := time.Tick(time.Second) //返回通道 每隔1S发送一条消息

	for {
		time.Sleep(1 * time.Second) //下游接受数据
		select {
		case n := <-c1:
			fmt.Printf("从c1接收到%d \n", n)
		case n := <-c2:
			fmt.Printf("从c2接收到%d \n", n)
		case <-tm:
			fmt.Println("bye")
			return
		case <-tick:
			fmt.Println("tick")
		default:
			fmt.Printf("No Value \n")
		}
	}

	/**---------------下游接受数据太慢 会跳过一些数据---------------------**/

}

func generator() chan int {
	out := make(chan int)

	go func() {
		i := 0
		for {
			//上游发送数据
			// 1500ms随机休眠
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}
