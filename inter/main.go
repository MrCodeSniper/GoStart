package main

import (
	"awesomeProject/impl"
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

func invoke(r Retriver) string { //这个方法接受一个类实现了GetContent方法 再调用这个类的GetContent方法
	return r.GetContent("http://www.baidu.com")
}

func main() {
	//实现类
	retriver := impl.Retriver{"this"}
	fmt.Print(invoke(retriver))
	fmt.Println()
	//connect:=impl.NetConnect{"陈鸿的PC",time.Minute}
	//fmt.Println(invoke(connect))

	var r Retriver
	r = impl.Retriver{} //
	//因为需要传入指针接受值
	r = impl.NetConnect{"陈鸿的PC", time.Minute} // 接口 可以为 一个实现了接口方法的构造类
	r = impl.NetConnect{
		UserAgent: "test",
		TimeOut:   time.Hour,
	}
	fmt.Printf("%T ,%v\n", r, r) //接口的具体类型  具体类型的成员变量

	//inspect(r) typeassertion类型判断

	types, ok := r.(impl.NetConnect)

	if ok {
		fmt.Println("xxx", types, types.TimeOut)
	}

}

func inspect(r Retriver) {
	switch v := r.(type) {
	case impl.NetConnect:
		fmt.Println(v.GetContent("http://www.baidu.com"))
	case impl.Retriver:
		fmt.Println("xxx")
	}
}
