package main

import (
	"fmt"
	"io"
	"math"
	"strings"
)

//函数getNumber 返回 返回值为int的函数
func getNumber()  func() int{
	i:=0
	//直接返回匿名函数
	return func() int{
		i+=1
		return i
	}
}



func totalAdd() func(int) int{
	sum:=0
	return func(i int) int {
		 sum= sum+i
		 return sum
	}
}

type intGeneration func() int //设置别名

//实现reader接口
func (gen intGeneration) Read(p []byte) (n int, err error) {
	next:=gen() //将相加的方法给next
	if(next>10000){
		return 0,io.EOF
	}
	s:=fmt.Sprintf("%d\n",next)//转换为字符串
	return strings.NewReader(s).Read(p)//将s写到字节数组中并返回
}




func fibonacci() intGeneration{
	a,b:=0,1
	 return func() int {
	 	a,b=b,a+b
	 	return a
	 }
}




func CompositeFunc(a int,f func(int) int) int{
	return f(a)
}


func main(){

	//自增函数
	adder:=getNumber() //函数为变量 直接()调用

	fmt.Println(adder())
	fmt.Println(adder())


	adder2:=getNumber() //创建新的函数 内部有对应的变量 保存
	fmt.Println(adder2())


    square:= func(x float64) float64{ //
    	return math.Sqrt(x)
	}

    fmt.Println(square(9))


    totalAdder:=totalAdd()

    for i:=0 ; i<10; i++{
    	fmt.Println(totalAdder(i))
	}


    //例 实现斐波那契数列  1 1 2 3 5 8
    fmt.Println("实现斐波那契数列")
    fab:=fibonacci()

    fmt.Println(fab())
	fmt.Println(fab())
	fmt.Println(fab())
	fmt.Println(fab())
	fmt.Println(fab())
	fmt.Println(fab())

    fmt.Println("函数操作")

   //函数为参数 将变量a进行函数操作
    num:= CompositeFunc(4, func(i int) int {
   	  i=i*i
   	  return i
   })





    fmt.Println(num)






}