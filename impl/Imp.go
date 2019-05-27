package impl

import (
	"net/http"
	"net/http/httputil"
	"time"
)

//类
type Retriver struct {
	Content string
}

//这个类实现了Get方法
func (retriver Retriver) GetContent(url string) string {
	return retriver.Content + url
}

//实现post方法
func (retriver Retriver)  Post(str string,form map[string]string)string{
	 retriver.Content=form["test"]
	 return "ok"
}










type NetConnect struct {
	UserAgent string        //用户设备
	TimeOut   time.Duration //超时时间
}

//将接口改成指针访问形式 因为struce很大 不想用copy 而是直接访问
//一个Get请求sample
func (connect NetConnect) GetContent(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	bytes, err := httputil.DumpResponse(resp, true)
	resp.Body.Close() //关闭资源

	if err != nil {
		panic(err)
	}

	return string(bytes)
}




//调用者
