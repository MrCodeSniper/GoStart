package main

import (
	"GoStart/Sirupsen/logrus"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	//http.HandleFunc("/list/", errorWrapper(fileList))
	////http://localhost:8887/list/fib.txt
	//http.ListenAndServe(":8887",nil)

	tryRecover()
}


func tryRecover(){


	defer func() {

		r:=recover()


		if err,ok := r.(error);ok{
			//如果这个类型是错误
			log1 := logrus.New()
			log1.Warn("错误信息",err)
		}else{
			panic(err)
		}

	}()

	//panic(errors.New("this is my error"))


	b:=0
     a:=5/b

     fmt.Println(a)


}


//4.包装错误处理
func errorWrapper(handler appHandler) func (writer http.ResponseWriter, request *http.Request){
	//将错误在这里返回 返回函数无需包含错误
	return func(writer http.ResponseWriter, request *http.Request) {
		//在函数执行时将error打出
		err:=handler(writer,request)
		if err!=nil{
			errCode:=http.StatusOK
			switch  {
			//如果错误是文件不存在
			case os.IsNotExist(err):
				errCode=http.StatusNotFound
			default:
				errCode=http.StatusInternalServerError
			}
			log1 := logrus.New()
			log1.Warn("错误信息",errCode)
			http.Error(writer,http.StatusText(errCode),errCode)//打印状态码给外界
		}
	}
}

//3.定义别名
type appHandler func (writer http.ResponseWriter, request *http.Request) error

//1.提出匿名函数 2.返回错误 交给外部处理
func fileList(writer http.ResponseWriter, request *http.Request) error{
	path:=request.URL.Path[len("/list/"):]//拿到字符串进行切除
	file, e := os.Open(path)
	if e!=nil{
		return e
	}

	defer file.Close()

	bytes, error := ioutil.ReadAll(file)//将文件读取拿到字节数组

	if error!=nil{
		return error
	}

	writer.Write(bytes)
	return nil
}
