package wei

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

)

type Context struct {
	Writer  http.ResponseWriter
	Req  *http.Request
}

func newContext(w http.ResponseWriter, req *http.Request) *Context  {
	return &Context{
		Writer: w,
		Req: req,
	}
}

// 接收Post请求的form参数(form-data, x-www-form-urlencoded)
func (c* Context) PostForm(key string) string {

	return c.Req.FormValue(key)
}

// 接收Get请求的query参数(url上的参数)
func (c *Context) GetQuery(key string) string {
	return c.Req.URL.Query().Get(key)
}

// 接收前端json数据
func (c *Context)JsonBody(v interface{}) {
	all, _ := ioutil.ReadAll(c.Req.Body)
	defer c.Req.Body.Close()


	err := json.Unmarshal(all, v)
	if err != nil{
		fmt.Println(err)
	}
	return
}

// 响应状态码
func (c *Context) StatusCode(code int)  {
	c.Writer.WriteHeader(code)
}

// 响应头添加参数
func (c *Context) SetHeader(key string, value string)  {
	c.Writer.Header().Set(key, value)
}

// 返回字符串参数
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.StatusCode(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// 返回Json
func (c *Context) Json(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application-json")
	c.StatusCode(code)
	marshal, _ := json.Marshal(obj)
	c.Writer.Write(marshal)
}

