package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	connect()
}


func connect(){

	var url string ="https://www.imooc.com/"

	//新建请求
	request, e := http.NewRequest(http.MethodGet, url, nil)

	if e!=nil {
		panic(e)
	}


	//增加头部
    request.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")


	//自建httpclient

	resp, err := http.DefaultClient.Do(request) //http客户端请求

	if err==nil{
		fmt.Println(resp)
	}else{
		panic(err)
	}

	defer resp.Body.Close() //关闭资源

	//打印返回数据为 字节数组
	bytes, err := httputil.DumpResponse(resp, true)

	if err!=nil{
		panic(err)
	}else{

		fmt.Printf("%s", bytes)
	}


}
