package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是主页")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是用户")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是创建用户")
}

func order(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是订单")
}

func main() {
	server := NewhttpServer("test-server")
	server.Route("/", home)
	server.Route("/user", user)
	server.Route("/user/create", createUser)
	server.Route("/order", order)
	server.Route("/user/signup", signup)
	server.Start(":8080", nil)
}

type Server interface {
	Route(pattern string, handlerFunc http.HandlerFunc)
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

type signUpReq struct {
	// 后面的tag 标签 运行期通过反射拿到？？申明式API会用这个？go arm框架好多用这个？
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}
type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// func signup(w http.ResponseWriter, r *http.Request) {
// 注册路由2
func signup(ctx *Context) {

	req := &signUpReq{}

	ctx := &Context{
		W: w,
		R: r,
	}
	err := ctx.ReadJson(req)
	if err != nil {
		fmt.Fprintf(w, "err:%v", err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		// 要返回掉，不然就会继续执行后面的代码
		return
	}
	err = json.Unmarshal(body, req)
	if err != nil {
		fmt.Fprintf(w, "deserialized failed: %v", err)
		// 要返回掉，不然就会继续执行后面的代码
		return
	}
	resp := &commonResponse{
		Data: 123,
	}

	w.WriteHeader(http.StatusOK)
	respJson, err := json.Marshal(resp)
	if err != nil {

	}

	// 返回一个虚拟的 user id 表示注册成功了 respJson 这里是byte类型，要转string 2者可以互转 func Marshal(v any) ([]byte, error) {
	fmt.Fprintf(w, string(respJson))
}
