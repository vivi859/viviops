package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Context struct {
	//这个是接口 ，不需要用*指针
	W http.ResponseWriter
	// Request是一个结构体所以要用指针
	R *http.Request
}

func (c *Context) ReadJson(v interface{}) error {
	//读body处理
	r := c.R
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		// 记住要返回，不然就还会执行后面的代码
		return err
	}
	err = json.Unmarshal(body, v)
	// 尝试再次读取，啥也读不到，但是也不会报错
	//body, err = io.ReadAll(r.Body)
	if err != nil {
		// 不会进来这里
		//fmt.Fprintf(w, "read the data one more time got error: %v", err)
		return err
	}
	//这里不需要return req ?
	return nil
}

// 接受任意类型的话就用interface{} 空接口
func (c *Context) WriteJson(code int, resp interface{}) error {
	c.W.WriteHeader(code)
	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	//err 前面申明过了就不需要再用:=了，
	_, err = c.W.Write(respJson)
	return err
}

func (c *Context) OkJson(resp interface{}) error {
	return c.WriteJson(http.StatusOK, resp)
}
