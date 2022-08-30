package wei

import (
	"fmt"
	"net/http"
)


type HandleFunc func(*Context)

type  Handle struct{
	method string
	handleFunc HandleFunc
}

// 路由
type Routes struct {
	middwa map[string][]HandleFunc
	route map[string]Handle
}

func New() *Routes {
	return &Routes{route: make(map[string]Handle),
					middwa: make(map[string][]HandleFunc)}
}


// 添加路由
func (rs *Routes) addroute(method string, requestpath string, handle HandleFunc)  {
	rs.route[requestpath] = Handle{
		method : method,
		handleFunc: handle,
	}
}

// 把中间件函数加入中间件
func (rs *Routes)addmidd(requestpath string, MiddleHandFun ...HandleFunc)  {
	if len(MiddleHandFun)  == 0 {
		return
	}
	for _, k:= range MiddleHandFun{
		rs.middwa[requestpath] = append(rs.middwa[requestpath], k)
	}

}

func (rs *Routes) Get(requestpath string, handle HandleFunc, MiddleHandFun ...HandleFunc)  {
	rs.addmidd(requestpath, MiddleHandFun...)
	rs.addroute("GET", requestpath, handle)
}

func (rs *Routes) Post(requestpath string, handle HandleFunc, MiddleHandFun ...HandleFunc)  {
	rs.addmidd(requestpath, MiddleHandFun...)
	rs.addroute("POST", requestpath, handle)
}

func (rs *Routes)Run(addr string)(err error)  {
	return http.ListenAndServe(addr, rs)
}

func (rs *Routes) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	fmt.Printf("%s>>>%s\n", req.Method, req.URL.Path)
	// 判断url是否注册
	if handle, ok := rs.route[req.URL.Path]; ok{
		if req.Method == handle.method{
			handle.handleFunc(c)
		}else {
			fmt.Fprintf(w, "404 请求方式错误!>>URL:%s>>Method: %s\n", req.URL, req.Method)
		}
	}else {

		fmt.Fprintf(w, "404 路由不存在>>URL:%s>>Method: %s\n", req.URL, req.Method)
	}

	// 执行中间件函数
	url := req.URL.Path
	if hf, ok := rs.middwa[url]; ok{
		for _, v := range hf{
			v(c)
		}
	}
}
