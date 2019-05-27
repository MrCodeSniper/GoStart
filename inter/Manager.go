package main

import (
	"GoStart/fib"
	"bufio"
	"errors"
	"fmt"
	"os"
)

func tryDefer(){
	//defer 关键词 在函数结束后打印 用栈管理 先进后出
	//如果在循环里defer 退出循环时倒序
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}



//文件操作
func writeFile(filename string){
	file, e := os.Create(filename)
	if e!=nil {
		panic(e)
	}
	defer file.Close()

	writer := bufio.NewWriter(file) //创建一个文件的缓冲输出流
	//需要将缓冲区里的数据刷出
	defer writer.Flush()

	funcc:=fib.Fibonacci()//获得斐波那契的函数

	for i:=0;i<20;i++{
		fmt.Fprint(writer,"|")
		fmt.Fprint(writer, funcc())//将函数执行的返回值循环写入writer
	}

}






func main() {
	tryDefer()
	writeFile("fib.txt")
	//错误处理 error是一个接口  类型为*os.PathEroor 分为op，path,Err 操作描述，操作路径，错误描述
    err:=errors.New("custom error")
    panic(err)
}
