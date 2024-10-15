package main

import (
	"encoding/json"
	"fmt"
	"go/parser"
	"net/http"
)

// 这个方法接口可以随便写到哪里，不需要引入包，就可以实现接口，在项目其他地方都可以实现？？？
type Server interface {
	//这里其实就是func 但是不用加func关键字，接口一般大写开头，全局使用，实现一般用小写。 接口是一组行为的抽象。
	//Route(pattern string, handlerFunc http.HandlerFunc)
	//注册路由2
	Route(pattern string, handlerFunc func(ctx *Context))


	//启动服务器
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

// 注册路由
//func (s *sdkHttpServer) Route(pattern string, handlerFunc http.HandlerFunc) {
//	http.HandleFunc(pattern, handlerFunc)
//}
//注册路由2
func(s * sdkHttpServer) Route(pattern string,handlerFunc func(ctx *Context)){
	http.HandleFunc(pattern,func(write http.ResponseWriter,r http.Request){
		ctx:=NewContext(write,request)
		handlerFunc(ctx)
	})
}

func NewContext(write http.ResponseWriter,request *http.Request) * Context{
	return
}

func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address,handler:nil)
}
func NewhttpServer(name string) Server{
	return &sdkHttpServer{//返回实际类型所实验的接口的时候，要加& 指针
		Name: name,
	}
}

//	func NewServer() Server(){
//		return &sdkHttpServer{}
//	}
//
// 返回server接口，factory是一个函数
type Factory func() Server

// 全局变量facory 类型是Factory
var facory Factory

func RegisterFactory(f Factory) {
	facory = f
}
func Newserver() Server {
	return facory()
}

func Signup(w http.ResponseWriter, r *http.Request) {
	req := &signUpReq{}
	ctx := &Context{
		W: w,
		R: r,
	}
	err := ctx.ReadJson(req)
	if err != nil {
		fmt.Fprintf("err:%v",err)
		return
	}


	resp:=&commonResponse{
		Data: 123,
	}
	respJson, err := json.Marshal(resp)
	if err!= nil{}
	//返回一个虚拟的user id 表示注册成功了
	fmt.Fprintf(w,string(respJson))

}
