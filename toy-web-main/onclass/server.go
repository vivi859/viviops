package main

import (
	"fmt"
	"net/http"
)
//这个方法接口可以随便写到哪里，不需要引入包，就可以实现接口，在项目其他地方都可以实现？？？
type Server interface {
	//这里其实就是func 但是不用加func关键字，接口一般大写开头，全局使用，实现一般用小写。 接口是一组行为的抽象。
	Route(pattern string, handlerFunc http.HandlerFunc)
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

//type Heander map[string][]string

//注册路由
func (s (sdkHttpServer) Route(pattern string, handlerFunc http.HandlerFunc) {
	panic("implement me")
})

func (s sdkHttpServer) Start(address string) error {
	panic("implement me")
}

func NewServer() Server(){
	return &sdkHttpServer{}
}

func NewServerWithName(name string) Server{

}

type Factory func() Server

func Signup(w http.ResponseWriter, r *http.Request){
	req := &SignupReq{}
	ctx := &Context{
		W: w,
		R: r,
	}
	err := ctx.ReadJson(req)
	if err != nil{
		 fmt.Fprintf("")
	}
}