package wei

import (
	"strings"
)

type Group struct {
	Rou *Routes
	Mainpath string
}

func NewGroup(h *Routes, mp string) *Group {
	return &Group{
		Rou: h,
		Mainpath: mp,
	}
}

func (g *Group)Get(requestpath string, handle HandleFunc, MiddleHandFun ...HandleFunc)  {
	var build strings.Builder
	build.WriteString(g.Mainpath)
	build.WriteString(requestpath)
	s3 := build.String()
	g.Rou.addmidd(s3, MiddleHandFun...)
	g.Rou.addroute("GET", s3, handle)

}

func (g *Group)Post(requestpath string, handle HandleFunc, MiddleHandFun ...HandleFunc)  {
	var build strings.Builder
	build.WriteString(g.Mainpath)
	build.WriteString(requestpath)
	s3 := build.String()
	g.Rou.addmidd(s3, MiddleHandFun...)
	g.Rou.addroute("POST", s3, handle)
}