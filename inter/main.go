package main

import (
	"GoStart/impl"
	"GoStart/queue"
	"fmt"
	"time"
)

//1.定义一个接口具备相应的行为
//2.定义多个类型 具备这个接口相同的行为 而不建立依赖或组合联系
//3.再需要一个接口的实现方法
//这个接口的实现方法 可以接受具备这样行为的结构体
type Retriver interface {
	GetContent(str string) string
}

type Poster interface { //定义一个接口Post
	Post(str string,form map[string]string)string
}

type Composite interface { //组合接口
	Retriver
	Poster
}

func invoke(r Retriver) string { //这个方法接受一个类实现了GetContent方法 再调用这个类的GetContent方法
	return r.GetContent("http://www.baidu.com")
}


//默认调用方法
func post(poster Poster)string{
	return poster.Post("xxx", map[string]string{
		"test":"ddd",
	})
}

func session(composite Composite) string{
	composite.Post("yyy", map[string]string{
		"test":"xxx",
	})
	return composite.GetContent("pppp")
}

func main() {
	//实现类
	retriver := &impl.Retriver{"this"}
	fmt.Print(invoke(retriver))
	fmt.Println()
	//connect:=impl.NetConnect{"陈鸿的PC",time.Minute}
	//fmt.Println(invoke(connect))

	var r Retriver
	r = &impl.Retriver{} //
	//因为需要传入指针接受值
	r = impl.NetConnect{"陈鸿的PC", time.Minute} // 接口 可以为 一个实现了接口方法的构造类
	r = impl.NetConnect{
		UserAgent: "test",
		TimeOut:   time.Hour,
	}
	fmt.Printf("%T ,%v\n", r, r) //接口的具体类型  具体类型的成员变量

	//inspect(r)

	// typeassertion 类型判断
	types, ok := r.(impl.NetConnect)
	if ok {
		fmt.Println("xxx", types, types.TimeOut)
	}



    q:=queue.Queue{}
    q.Push("字符串")
    q.Push(5)
    q.Push(true)
    q.Pop()



    fmt.Println(q,q.IsEmpty())


    fmt.Println(session(retriver))



	//go语言自带 唯一满足并发条件的类型 通道

	//初始化


	//chan int //元素类型为int的通道

	//根据通道容量的大小称为 非缓冲通道和缓冲通道

	//一个通道相当于一个先进先出（FIFO）的队列
	//操作符 -> <-

	ch := make(chan int,3) //通道容量为3
	ch <- 2 //入队
	ch <- 1
	ch <- 3
	elem:=<-ch//出队

	fmt.Println(elem)

	//作为通道来说 视角限定
	ch1 :=make(chan<- int,3) //创建单向的接受通道
	ch2 :=make(<-chan int,3) //创建单向的发送通道

	fmt.Println(ch1,ch2) //打印变量的内存地址



	c:=make(chan int)

	//fatal error: all goroutines are asleep - deadlock!
	//channel 是goroutine之间交互的媒介 不能在相同的go routine中收 需要新建goroutine
	//必须先开协程再入通道
	//c<-1
	//c<-2
	//n:= <-c
	//fmt.Println(n)

	// fatal error: all goroutines are asleep - deadlock! 通道没有数据 取数据导致堵塞
	//go func() {
	//	v:= <-c
	//	fmt.Println(v)
	//}()





	//这里只打印了1个值 因为子线程执行到第二步时 进程已经结束了 需要主线程wait
	go func() {
		for{
			v:= <-c
			fmt.Println(v)
		}
	}()


	c<- 3
	c<- 4

	time.Sleep(time.Second)


	go Workder(c) //通道作为参数 子线程函数


	c<- 5
	c<- 6



	time.Sleep(time.Second)


	c2:=createWorker() //返回channel 内部包含channel操作

	c2<- 0
	c2<- 10

	close(c2)//关闭通道 之后接收到的值都为0 需要判断并退出死循环

	time.Sleep(time.Second)

	//作用为限定类型  接口方法接受单向通道 比如发送方法  只能使用接受通道

//对通道的发送和接收操作都有哪些基本的特性
/**

	发送方接收方共享缓冲 而socket是接收发送方都有缓冲

  对于同一个通道，发送操作之间是互斥的，接收操作之间也是互斥的。
  （也就是说通道的操作是同步的不能并发 未完成前会堵塞 ）


  发送操作和接收操作中对元素值的处理都是不可分割的。（一次操作是原子性的 要么失败要么成功）
  发送操作在完全完成之前会被阻塞。接收操作也是如此。
  通道也是值传递 进入的是参数的副本  1.先生成副本2.防止副本到通道内部3.删除原值

  通道的长度就是通道的元素个数 容量即初始化的最大数量

  特殊情况：
  缓冲通道 如果通道满 除非有元素出去 不然一直堵着
  同理如果通道空 出去的操作会一直堵塞

  非缓冲通道适合 例如微信交流的场景 用户获取信息并展示 必须有信息来

  通道关闭 再次操作会引发panic


   底层：

   通道底层为环形链表





 */











}

//chan 作为参数
func Workder(c chan int){
	for{
		v:= <-c
		fmt.Println(v)
	}
}

//chan 作为返回值

func createWorker() chan int{
	c:=make(chan int)
	go func() {
		for{
			v,ok:=<-c
			if !ok{ //没有值退出循环 不再接收值 或者通过range 缓冲容量 接收固定个数值
				break
			}
			fmt.Println(v)
		}
	}()
	return c
}


//关闭channel






//typeswitch 类型筛选
func inspect(r Retriver) {
	switch v := r.(type) {
	case impl.NetConnect:
		fmt.Println(v.GetContent("http://www.baidu.com"))
	case *impl.Retriver:
		fmt.Println("xxx")
	}
}
